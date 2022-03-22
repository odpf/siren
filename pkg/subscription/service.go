package subscription

import (
	"context"
	"fmt"
	"sort"

	"github.com/odpf/siren/domain"
	"github.com/odpf/siren/pkg/namespace"
	"github.com/odpf/siren/pkg/provider"
	"github.com/odpf/siren/pkg/receiver"
	"github.com/odpf/siren/pkg/subscription/alertmanager"
	"github.com/odpf/siren/store"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Service handles business logic
type Service struct {
	repository       store.SubscriptionRepository
	providerService  domain.ProviderService
	namespaceService domain.NamespaceService
	receiverService  domain.ReceiverService
	amClient         alertmanager.Client
}

// NewService returns service struct
func NewService(repository store.SubscriptionRepository, providerRepository store.ProviderRepository, namespaceRepository store.NamespaceRepository,
	receiverRepository store.ReceiverRepository, db *gorm.DB, key string) (domain.SubscriptionService, error) {
	namespaceService, err := namespace.NewService(namespaceRepository, key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create namespace service")
	}
	receiverService, err := receiver.NewService(receiverRepository, nil, key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create receiver service")
	}
	return &Service{repository, provider.NewService(providerRepository),
		namespaceService, receiverService, nil}, nil
}

func (s Service) ListSubscriptions(ctx context.Context) ([]*domain.Subscription, error) {
	subscriptions, err := s.repository.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "s.repository.List")
	}

	return subscriptions, nil
}

func (s Service) CreateSubscription(ctx context.Context, sub *domain.Subscription) (*domain.Subscription, error) {
	ctx = s.repository.WithTransaction(ctx)
	sortReceivers(sub)
	newSubscription, err := s.repository.Create(ctx, sub)
	if err != nil {
		if err := s.repository.Rollback(ctx); err != nil {
			return nil, errors.Wrap(err, "s.repository.Rollback")
		}
		return nil, errors.Wrap(err, "s.repository.Create")
	}

	if err := s.syncInUpstreamCurrentSubscriptionsOfNamespace(ctx, newSubscription.Namespace); err != nil {
		if err := s.repository.Rollback(ctx); err != nil {
			return nil, errors.Wrap(err, "s.repository.Rollback")
		}
		return nil, errors.Wrap(err, "s.syncInUpstreamCurrentSubscriptionsOfNamespace")
	}

	if err := s.repository.Commit(ctx); err != nil {
		return nil, errors.Wrap(err, "s.repository.Commit")
	}
	return newSubscription, nil
}

func (s Service) GetSubscription(ctx context.Context, id uint64) (*domain.Subscription, error) {
	subscription, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "s.repository.Get")
	}
	if subscription == nil {
		return nil, nil
	}
	return subscription, nil
}

func (s Service) UpdateSubscription(ctx context.Context, sub *domain.Subscription) (*domain.Subscription, error) {
	ctx = s.repository.WithTransaction(ctx)
	sortReceivers(sub)
	updatedSubscription, err := s.repository.Update(ctx, sub)
	if err != nil {
		if err := s.repository.Rollback(ctx); err != nil {
			return nil, errors.Wrap(err, "s.repository.Rollback")
		}
		return nil, errors.Wrap(err, "s.repository.Update")
	}

	if err := s.syncInUpstreamCurrentSubscriptionsOfNamespace(ctx, updatedSubscription.Namespace); err != nil {
		if err := s.repository.Rollback(ctx); err != nil {
			return nil, errors.Wrap(err, "s.repository.Rollback")
		}
		return nil, errors.Wrap(err, "s.syncInUpstreamCurrentSubscriptionsOfNamespace")
	}

	if err := s.repository.Commit(ctx); err != nil {
		return nil, errors.Wrap(err, "s.repository.Commit")
	}
	return updatedSubscription, nil
}

func (s Service) DeleteSubscription(ctx context.Context, id uint64) error {
	sub, err := s.repository.Get(ctx, id)
	if err != nil {
		return errors.Wrap(err, "s.repository.Get")
	}

	ctx = s.repository.WithTransaction(ctx)
	if err := s.repository.Delete(ctx, id); err != nil {
		if err := s.repository.Rollback(ctx); err != nil {
			return errors.Wrap(err, "s.repository.Rollback")
		}
		return errors.Wrap(err, "s.repository.Delete")
	}

	if err := s.syncInUpstreamCurrentSubscriptionsOfNamespace(ctx, sub.Namespace); err != nil {
		if err := s.repository.Rollback(ctx); err != nil {
			return errors.Wrap(err, "s.repository.Rollback")
		}
		return errors.Wrap(err, "s.syncInUpstreamCurrentSubscriptionsOfNamespace")
	}
	return nil
}

func (s Service) Migrate() error {
	return s.repository.Migrate()
}

var alertmanagerClientCreator = alertmanager.NewClient

func (s Service) syncInUpstreamCurrentSubscriptionsOfNamespace(ctx context.Context, namespaceId uint64) error {
	// fetch all subscriptions in this namespace.
	subscriptionsInNamespace, err := s.getAllSubscriptionsWithinNamespace(ctx, namespaceId)
	if err != nil {
		return errors.Wrap(err, "s.getAllSubscriptionsWithinNamespace")
	}
	// check provider type of the namespace
	providerInfo, namespaceInfo, err := s.getProviderAndNamespaceInfoFromNamespaceId(namespaceId)
	if err != nil {
		return errors.Wrap(err, "s.getProviderAndNamespaceInfoFromNamespaceId")
	}
	subscriptionsInNamespaceEnrichedWithReceivers, err := s.addReceiversConfiguration(subscriptionsInNamespace)
	if err != nil {
		return errors.Wrap(err, "s.addReceiversConfiguration")
	}
	// do upstream call to create subscriptions as per provider type
	switch providerInfo.Type {
	case "cortex":
		amConfig := getAmConfigFromSubscriptions(subscriptionsInNamespaceEnrichedWithReceivers)
		newAMClient, err := alertmanagerClientCreator(domain.CortexConfig{Address: providerInfo.Host})
		if err != nil {
			return errors.Wrap(err, "alertmanagerClientCreator: ")
		}
		s.amClient = newAMClient
		err = s.amClient.SyncConfig(amConfig, namespaceInfo.Urn)
		if err != nil {
			return errors.Wrap(err, "s.amClient.SyncConfig")
		}
	default:
		return errors.New(fmt.Sprintf("subscriptions for provider type '%s' not supported", providerInfo.Type))
	}
	return nil
}

func (s Service) getAllSubscriptionsWithinNamespace(ctx context.Context, id uint64) ([]*domain.Subscription, error) {
	subscriptions, err := s.repository.List(ctx) // TODO: pass namespaceID as list filter
	if err != nil {
		return nil, errors.Wrap(err, "s.repository.List")
	}
	var subscriptionsWithinNamespace []*domain.Subscription
	for _, sub := range subscriptions {
		if sub.Namespace == id {
			subscriptionsWithinNamespace = append(subscriptionsWithinNamespace, sub)
		}
	}
	return subscriptionsWithinNamespace, nil
}

func (s Service) getProviderAndNamespaceInfoFromNamespaceId(id uint64) (*domain.Provider, *domain.Namespace, error) {
	namespaceInfo, err := s.namespaceService.GetNamespace(id)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get namespace details")
	}
	providerInfo, err := s.providerService.GetProvider(namespaceInfo.Provider)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get provider details")
	}
	return providerInfo, namespaceInfo, nil
}

func (s Service) addReceiversConfiguration(subscriptions []*domain.Subscription) ([]SubscriptionEnrichedWithReceivers, error) {
	res := make([]SubscriptionEnrichedWithReceivers, 0)
	allReceivers, err := s.receiverService.ListReceivers()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get receivers")
	}
	for _, item := range subscriptions {
		enrichedReceivers := make([]EnrichedReceiverMetadata, 0)
		for _, receiverItem := range item.Receivers {
			var receiverInfo *domain.Receiver
			for idx := range allReceivers {
				if allReceivers[idx].Id == receiverItem.Id {
					receiverInfo = allReceivers[idx]
					break
				}
			}
			if receiverInfo == nil {
				return nil, errors.New(fmt.Sprintf("receiver id %d does not exist", receiverItem.Id))
			}
			//initialize the nil map using the make function
			//to avoid panics while adding elements in future
			if receiverItem.Configuration == nil {
				receiverItem.Configuration = make(map[string]string)
			}
			switch receiverInfo.Type {
			case "slack":
				if _, ok := receiverItem.Configuration["channel_name"]; !ok {
					return nil, errors.New(fmt.Sprintf(
						"configuration.channel_name missing from receiver with id %d", receiverItem.Id))
				}
				if val, ok := receiverInfo.Configurations["token"]; ok {
					receiverItem.Configuration["token"] = val.(string)
				}
			case "pagerduty":
				if val, ok := receiverInfo.Configurations["service_key"]; ok {
					receiverItem.Configuration["service_key"] = val.(string)
				}
			case "http":
				if val, ok := receiverInfo.Configurations["url"]; ok {
					receiverItem.Configuration["url"] = val.(string)
				}
			default:
				return nil, errors.New(fmt.Sprintf(`subscriptions for receiver type %s not supported via Siren inside Cortex`, receiverInfo.Type))
			}
			enrichedReceiver := EnrichedReceiverMetadata{
				Id:            receiverItem.Id,
				Configuration: receiverItem.Configuration,
				Type:          receiverInfo.Type,
			}
			enrichedReceivers = append(enrichedReceivers, enrichedReceiver)
		}
		enrichedSubscription := SubscriptionEnrichedWithReceivers{
			Id:          item.Id,
			NamespaceId: item.Namespace,
			Urn:         item.Urn,
			Receiver:    enrichedReceivers,
			Match:       item.Match,
		}
		res = append(res, enrichedSubscription)
	}
	return res, nil
}

func sortReceivers(sub *domain.Subscription) {
	sort.Slice(sub.Receivers, func(i, j int) bool {
		return sub.Receivers[i].Id < sub.Receivers[j].Id
	})
}
