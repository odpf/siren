package service

import (
	"github.com/odpf/siren/pkg/namespace"
	"github.com/odpf/siren/pkg/provider"
	"github.com/odpf/siren/pkg/receiver"
	"github.com/odpf/siren/pkg/subscription"
	"github.com/odpf/siren/store"
	"net/http"

	"github.com/odpf/siren/domain"
	"github.com/odpf/siren/pkg/alerts"
	"github.com/odpf/siren/pkg/rules"
	"github.com/odpf/siren/pkg/slacknotifier"
	"github.com/odpf/siren/pkg/templates"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Container struct {
	TemplatesService    domain.TemplatesService
	RulesService        domain.RuleService
	AlertService        domain.AlertService
	NotifierServices    domain.NotifierServices
	ProviderService     domain.ProviderService
	NamespaceService    domain.NamespaceService
	ReceiverService     domain.ReceiverService
	SubscriptionService domain.SubscriptionService
}

func Init(repositories *store.RepositoryContainer, db *gorm.DB, c *domain.Config, httpClient *http.Client) (*Container, error) {
	templatesService := templates.NewService(repositories.TemplatesRepository)
	rulesService := rules.NewService(repositories.TemplatesRepository, db)
	alertHistoryService := alerts.NewService(db)

	slackNotifierService := slacknotifier.NewService()
	providerService := provider.NewService(repositories.ProviderRepository)
	namespaceService, err := namespace.NewService(repositories.NamespaceRepository, c.EncryptionKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create namespace service")
	}
	receiverService, err := receiver.NewService(repositories.ReceiverRepository, httpClient, c.EncryptionKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create receiver service")
	}
	subscriptionService, err := subscription.NewService(repositories.ProviderRepository,
		repositories.NamespaceRepository, repositories.ReceiverRepository, db, c.EncryptionKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create subscriptions service")
	}

	return &Container{
		TemplatesService: templatesService,
		RulesService:     rulesService,
		AlertService:     alertHistoryService,
		NotifierServices: domain.NotifierServices{
			Slack: slackNotifierService,
		},
		ProviderService:     providerService,
		NamespaceService:    namespaceService,
		ReceiverService:     receiverService,
		SubscriptionService: subscriptionService,
	}, nil
}

func (container *Container) MigrateAll(db *gorm.DB) error {
	err := container.TemplatesService.Migrate()
	if err != nil {
		return err
	}
	err = container.RulesService.Migrate()
	if err != nil {
		return err
	}
	err = container.AlertService.Migrate()
	if err != nil {
		return err
	}
	err = container.ProviderService.Migrate()
	if err != nil {
		return err
	}
	err = container.NamespaceService.Migrate()
	if err != nil {
		return err
	}
	err = container.ReceiverService.Migrate()
	if err != nil {
		return err
	}
	err = container.SubscriptionService.Migrate()
	if err != nil {
		return err
	}
	return nil
}
