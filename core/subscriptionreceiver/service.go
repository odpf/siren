package subscriptionreceiver

import (
	"context"
)

type Transactor interface {
	WithTransaction(ctx context.Context) context.Context
	Rollback(ctx context.Context, err error) error
	Commit(ctx context.Context) error
}

type Repository interface {
	List(context.Context, Filter) ([]Relation, error)
	BulkCreate(context.Context, []Relation) error
	BulkUpsert(context.Context, []Relation) error
	BulkSoftDelete(ctx context.Context, flt DeleteFilter) error
	BulkDelete(ctx context.Context, flt DeleteFilter) error
}

// Service handles business logic
type Service struct {
	repository Repository
}

// NewService returns service struct
func NewService(
	repository Repository,
) *Service {
	svc := &Service{
		repository: repository,
	}

	return svc
}

func (s *Service) List(ctx context.Context, flt Filter) ([]Relation, error) {
	relations, err := s.repository.List(ctx, flt)
	if err != nil {
		return nil, err
	}

	return relations, nil
}

func (s *Service) BulkCreate(ctx context.Context, rels []Relation) error {
	if err := s.repository.BulkCreate(ctx, rels); err != nil {
		return err
	}

	return nil
}

func (s *Service) BulkUpsert(ctx context.Context, rels []Relation) error {
	if err := s.repository.BulkUpsert(ctx, rels); err != nil {
		return err
	}

	return nil
}

func (s *Service) BulkSoftDelete(ctx context.Context, flt DeleteFilter) error {
	if err := s.repository.BulkSoftDelete(ctx, flt); err != nil {
		return err
	}

	return nil
}

func (s *Service) BulkDelete(ctx context.Context, flt DeleteFilter) error {
	if err := s.repository.BulkDelete(ctx, flt); err != nil {
		return err
	}

	return nil
}
