package provider

import (
	"context"

	"github.com/raystack/siren/pkg/errors"
)

// Service handles business logic
type Service struct {
	repository Repository
}

// NewService returns repository struct
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) List(ctx context.Context, flt Filter) ([]Provider, error) {
	return s.repository.List(ctx, flt)
}

func (s *Service) Create(ctx context.Context, prov *Provider) error {
	if prov == nil {
		return errors.ErrInvalid.WithMsgf("provider is nil")
	}

	err := s.repository.Create(ctx, prov)
	if err != nil {
		if errors.Is(err, ErrDuplicate) {
			return errors.ErrConflict.WithMsgf(err.Error())
		}
		return err
	}

	return nil
}

func (s *Service) Get(ctx context.Context, id uint64) (*Provider, error) {
	prov, err := s.repository.Get(ctx, id)
	if err != nil {
		if errors.As(err, new(NotFoundError)) {
			return nil, errors.ErrNotFound.WithMsgf(err.Error())
		}
		return nil, err
	}
	return prov, nil
}

func (s *Service) Update(ctx context.Context, prov *Provider) error {
	if prov == nil {
		return errors.ErrInvalid.WithMsgf("provider is nil")
	}

	err := s.repository.Update(ctx, prov)
	if err != nil {
		if errors.Is(err, ErrDuplicate) {
			return errors.ErrConflict.WithMsgf(err.Error())
		}
		if errors.As(err, new(NotFoundError)) {
			return errors.ErrNotFound.WithMsgf(err.Error())
		}
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, id)
}
