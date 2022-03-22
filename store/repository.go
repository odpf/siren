package store

import (
	"context"

	"github.com/odpf/siren/domain"
	"github.com/odpf/siren/store/model"
	"github.com/odpf/siren/store/postgres"
	"gorm.io/gorm"
)

type NamespaceRepository interface {
	Migrate() error
	List() ([]*model.Namespace, error)
	Create(*model.Namespace) (*model.Namespace, error)
	Get(uint64) (*model.Namespace, error)
	Update(*model.Namespace) (*model.Namespace, error)
	Delete(uint64) error
}

type ProviderRepository interface {
	Migrate() error
	List(map[string]interface{}) ([]*domain.Provider, error)
	Create(*domain.Provider) (*domain.Provider, error)
	Get(uint64) (*domain.Provider, error)
	Update(*domain.Provider) (*domain.Provider, error)
	Delete(uint64) error
}

type ReceiverRepository interface {
	Migrate() error
	List() ([]*model.Receiver, error)
	Create(*model.Receiver) (*model.Receiver, error)
	Get(uint64) (*model.Receiver, error)
	Update(*model.Receiver) (*model.Receiver, error)
	Delete(uint64) error
}

type TemplatesRepository interface {
	Upsert(*model.Template) (*model.Template, error)
	Index(string) ([]model.Template, error)
	GetByName(string) (*model.Template, error)
	Delete(string) error
	Render(string, map[string]string) (string, error)
	Migrate() error
}

type SubscriptionRepository interface {
	Transactor
	Migrate() error
	List(context.Context) ([]*domain.Subscription, error)
	Create(context.Context, *domain.Subscription) (*domain.Subscription, error)
	Get(context.Context, uint64) (*domain.Subscription, error)
	Update(context.Context, *domain.Subscription) (*domain.Subscription, error)
	Delete(context.Context, uint64) error
}

type Transactor interface {
	WithTransaction(ctx context.Context) context.Context
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type RepositoryContainer struct {
	ProviderRepository     ProviderRepository
	NamespaceRepository    NamespaceRepository
	TemplatesRepository    TemplatesRepository
	ReceiverRepository     ReceiverRepository
	SubscriptionRepository SubscriptionRepository
}

func NewRepositoryContainer(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		NamespaceRepository:    postgres.NewNamespaceRepository(db),
		ProviderRepository:     postgres.NewProviderRepository(db),
		ReceiverRepository:     postgres.NewReceiverRepository(db),
		TemplatesRepository:    postgres.NewTemplateRepository(db),
		SubscriptionRepository: postgres.NewSubscriptionRepository(db),
	}
}
