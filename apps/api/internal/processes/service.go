package processes

import (
	"context"
	"partiq/internal/shared"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

type Service interface {
	List(ctx context.Context) ([]Process, error)
	GetByID(ctx context.Context, id int64) (Process, error)
}

type service struct {
	repo Repository
}

var processCache = shared.NewCache(5 * time.Minute)

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context) ([]Process, error) {
	if val, ok := processCache.Get("processes:all"); ok {
		if cached, ok := val.([]Process); ok {
			log.Ctx(ctx).Info().Str("source", "cache").Msg("Returned all processes")
			return cached, nil
		}
	}

	log.Ctx(ctx).Info().Str("source", "db").Msg("Fetching all processes from DB")
	processes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	processCache.Set("processes:all", processes)
	return processes, nil
}

func (s *service) GetByID(ctx context.Context, id int64) (Process, error) {
	key := "processes:" + strconv.FormatInt(id, 10)

	if val, ok := processCache.Get(key); ok {
		if cached, ok := val.(Process); ok {
			log.Ctx(ctx).Info().Int64("id", id).Str("source", "cache").Msg("Returned process by ID")
			return cached, nil
		}
	}

	log.Ctx(ctx).Info().Int64("id", id).Str("source", "db").Msg("Fetching process by ID from DB")
	process, err := s.repo.GetProcessByID(ctx, int(id))
	if err != nil {
		return Process{}, err
	}

	processCache.Set(key, process)
	return process, nil
}
