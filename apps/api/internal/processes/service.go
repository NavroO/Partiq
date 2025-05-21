package processes

import "context"

type Service interface {
	List(ctx context.Context) ([]Process, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context) ([]Process, error) {
	return s.repo.GetAll(ctx)
}
