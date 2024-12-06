package medication

import (
	"context"

	"github.com/aborilov/hippo/business/medication/model"
	"github.com/google/uuid"
)

type service struct {
	repo model.Repository
}

func NewService(repo model.Repository) (model.Service, error) {
	svc := &service{
		repo: repo,
	}
	return svc, nil
}

func (s *service) Create(ctx context.Context, m *model.Medication) (*model.Medication, error) {
	m.ID = uuid.New()
	return s.repo.Create(ctx, m)
}

func (s *service) List(ctx context.Context) ([]*model.Medication, error) {
	return s.repo.List(ctx)
}

func (s *service) Get(ctx context.Context, id uuid.UUID) (*model.Medication, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) Update(ctx context.Context, m *model.Medication) (*model.Medication, error) {
	return s.repo.Update(ctx, m)
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
