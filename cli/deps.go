package cli

import (
	"context"
	"fmt"

	saltlog "github.com/goto/salt/log"
	"github.com/goto/siren/config"
	"github.com/goto/siren/core/alert"
	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/namespace"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/provider"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/rule"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/core/subscriptionreceiver"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/internal/store/postgres"
	"github.com/goto/siren/pkg/pgc"
	"github.com/goto/siren/pkg/secret"
	"github.com/goto/siren/plugins/providers"
	"github.com/goto/siren/plugins/receivers/file"
	"github.com/goto/siren/plugins/receivers/httpreceiver"
	"github.com/goto/siren/plugins/receivers/pagerduty"
	"github.com/goto/siren/plugins/receivers/slack"
	"github.com/goto/siren/plugins/receivers/slackchannel"
)

func InitDeps(
	ctx context.Context,
	logger saltlog.Logger,
	cfg config.Config,
	queue notification.Queuer,
	withProviderPlugin bool,
) (*api.Deps, *pgc.Client, map[string]notification.Notifier, *providers.PluginManager, error) {

	pgClient, err := pgc.NewClient(logger, cfg.DB)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	encryptor, err := secret.New(cfg.Service.EncryptionKey)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("cannot initialize encryptor: %w", err)
	}

	templateRepository := postgres.NewTemplateRepository(pgClient)
	templateService := template.NewService(templateRepository)

	logRepository := postgres.NewLogRepository(pgClient)
	logService := log.NewService(logRepository)

	// TODO: need to figure out the way on how to nicely load deps without plugins dependency
	// in case the caller does not need that
	var providerService api.ProviderService
	var configSyncers = make(map[string]namespace.ConfigSyncer, 0)
	var alertTransformers = make(map[string]alert.AlertTransformer, 0)
	var ruleUploaders = make(map[string]rule.RuleUploader, 0)
	var providersPluginManager *providers.PluginManager

	if withProviderPlugin {
		providersPluginManager = providers.NewPluginManager(logger, cfg.Providers)
		providerPluginClients := providersPluginManager.InitClients()
		providerPlugins, err := providersPluginManager.DispenseClients(providerPluginClients)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		if err := providersPluginManager.InitConfigs(ctx, providerPlugins, cfg.Log.Level); err != nil {
			return nil, nil, nil, nil, err
		}

		var supportedProviderTypes = []string{}
		for typ := range providerPlugins {
			supportedProviderTypes = append(supportedProviderTypes, typ)
		}
		providerRepository := postgres.NewProviderRepository(pgClient)
		providerService = provider.NewService(providerRepository, supportedProviderTypes)

		if len(providerPlugins) == 0 {
			logger.Warn("no provider plugins found!")
		}

		for k, pc := range providerPlugins {
			alertTransformers[k] = pc.(alert.AlertTransformer)
			configSyncers[k] = pc.(namespace.ConfigSyncer)
			ruleUploaders[k] = pc.(rule.RuleUploader)
		}
	}

	namespaceRepository := postgres.NewNamespaceRepository(pgClient)
	namespaceService := namespace.NewService(encryptor, namespaceRepository, providerService, configSyncers)

	ruleRepository := postgres.NewRuleRepository(pgClient)
	ruleService := rule.NewService(
		ruleRepository,
		templateService,
		namespaceService,
		ruleUploaders,
	)

	silenceRepository := postgres.NewSilenceRepository(pgClient)
	silenceService := silence.NewService(silenceRepository)

	// plugin receiver services
	slackPluginService := slack.NewPluginService(cfg.Receivers.Slack, encryptor)
	slackChannelPluginService := slackchannel.NewPluginService(cfg.Receivers.Slack, encryptor)
	pagerDutyPluginService := pagerduty.NewPluginService(cfg.Receivers.Pagerduty)
	httpreceiverPluginService := httpreceiver.NewPluginService(logger, cfg.Receivers.HTTPReceiver)
	filePluginService := file.NewPluginService()

	receiverRepository := postgres.NewReceiverRepository(pgClient)
	receiverService := receiver.NewService(
		receiverRepository,
		map[string]receiver.ConfigResolver{
			receiver.TypeSlack:        slackPluginService,
			receiver.TypeSlackChannel: slackChannelPluginService,
			receiver.TypeHTTP:         httpreceiverPluginService,
			receiver.TypePagerDuty:    pagerDutyPluginService,
			receiver.TypeFile:         filePluginService,
		},
	)

	subscriptionReceiverRepository := postgres.NewSubscriptionReceiverRepository(pgClient)
	subscriptionReceiverService := subscriptionreceiver.NewService(subscriptionReceiverRepository)

	subscriptionRepository := postgres.NewSubscriptionRepository(pgClient)
	subscriptionService := subscription.NewService(
		subscriptionRepository,
		logService,
		namespaceService,
		receiverService,
		subscriptionReceiverService,
	)

	// notification
	idempotencyRepository := postgres.NewIdempotencyRepository(pgClient)
	notificationRepository := postgres.NewNotificationRepository(pgClient)
	alertRepository := postgres.NewAlertRepository(pgClient)

	notificationDeps := notification.Deps{
		Logger:                logger,
		Cfg:                   cfg.Notification,
		Repository:            notificationRepository,
		Q:                     queue,
		LogService:            logService,
		IdempotencyRepository: idempotencyRepository,
		AlertRepository:       alertRepository,
		ReceiverService:       receiverService,
		SubscriptionService:   subscriptionService,
		SilenceService:        silenceService,
		TemplateService:       templateService,
	}

	notifierRegistry := map[string]notification.Notifier{
		receiver.TypeSlack:        slackPluginService,
		receiver.TypeSlackChannel: slackChannelPluginService,
		receiver.TypePagerDuty:    pagerDutyPluginService,
		receiver.TypeHTTP:         httpreceiverPluginService,
		receiver.TypeFile:         filePluginService,
	}

	routerRegistry := map[string]notification.Router{
		notification.RouterReceiver: notification.NewRouterReceiverService(
			notificationDeps,
			notifierRegistry,
		),
		notification.RouterSubscriber: notification.NewRouterSubscriberService(
			notificationDeps,
			notifierRegistry,
		),
	}

	dispatchServiceRegistry := map[string]notification.Dispatcher{
		notification.DispatchKindBulkNotification: notification.NewDispatchBulkNotificationService(
			notificationDeps,
			notifierRegistry,
			routerRegistry,
		),
		notification.DispatchKindSingleNotification: notification.NewDispatchSingleNotificationService(
			notificationDeps,
			notifierRegistry,
			routerRegistry,
		),
	}

	notificationService := notification.NewService(
		notificationDeps,
		dispatchServiceRegistry,
	)

	alertService := alert.NewService(
		cfg.Alert,
		logger,
		alertRepository,
		logService,
		notificationService,
		alertTransformers,
	)

	return &api.Deps{
		TemplateService:             templateService,
		RuleService:                 ruleService,
		AlertService:                alertService,
		ProviderService:             providerService,
		NamespaceService:            namespaceService,
		ReceiverService:             receiverService,
		SubscriptionService:         subscriptionService,
		SubscriptionReceiverService: subscriptionReceiverService,
		NotificationService:         notificationService,
		SilenceService:              silenceService,
	}, pgClient, notifierRegistry, providersPluginManager, nil
}
