package processes

import "context"

type Service interface {
	List(ctx context.Context) ([]Process, error)
	GetByID(ctx context.Context, id int64) (Process, error)
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

func (s *service) GetByID(ctx context.Context, id int64) (Process, error) {
	return s.repo.GetProcessByID(ctx, int(id))
}
