package proposals

import "context"

type Service interface {
	List(ctx context.Context) ([]Proposal, error)
	ListByProcessID(ctx context.Context, processID int) ([]Proposal, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context) ([]Proposal, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) ListByProcessID(ctx context.Context, processID int) ([]Proposal, error) {
	return s.repo.GetByProcessID(ctx, processID)
}
