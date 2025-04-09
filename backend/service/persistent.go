package service

import (
	"context"
	"geoindexing_comparison/backend/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"time"
)

type RunRequest struct {
	Indexes []string `json:"indexes"`
	Tasks   []string `json:"tasks"`
	Start   uint64   `json:"start"`
	Stop    uint64   `json:"stop"`
	Step    uint64   `json:"step"`
}

func (r *Service) AddRun(ctx context.Context, req *RunRequest, createdBy string) (*repository.Run, error) {
	run := &repository.Run{
		CreatedAt: time.Now().UTC(),
		CreatedBy: createdBy,
		Status:    repository.RunStatusPending,
		Start:     req.Start,
		Stop:      req.Stop,
		Step:      req.Step,
		Indexes:   req.Indexes,
		Tasks:     req.Tasks,
	}

	err := r.repository.SaveRun(ctx, run)
	if err != nil {
		return nil, errors.Wrap(err, "could not save run")
	}

	zerolog.Ctx(ctx).Info().Interface("run", run).Msg("run.added")

	return run, nil
}

func (r *Service) StopRun(ctx context.Context, id int) (*repository.Run, error) {
	run, err := r.repository.GetRun(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not get run")
	}

	run.Status = repository.RunStatusCancelled

	err = r.repository.SaveRun(ctx, run)
	if err != nil {
		return nil, errors.Wrap(err, "could not save run")
	}

	zerolog.Ctx(ctx).Info().Interface("run", run).Msg("run.stopped")

	return run, nil
}

func (r *Service) GetRuns(ctx context.Context) ([]repository.Run, error) {
	return r.repository.GetRuns(ctx)
}
