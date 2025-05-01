package service

import (
	"context"
	"geoindexing_comparison/backend/repository"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
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

func (r *Service) StopRuns(ctx context.Context) error {
	err := r.repository.StopRuns(ctx)
	if err != nil {
		return errors.Wrap(err, "could not save run")
	}

	zerolog.Ctx(ctx).Info().Msg("runs.stopped")

	return nil
}

func (r *Service) GetRuns(ctx context.Context) ([]repository.Run, error) {
	return r.repository.GetRuns(ctx)
}

type Point struct {
	RunID uint64 `json:"runId"`
	Idx   uint64 `json:"idx"`

	Index string  `json:"index"`
	Task  string  `json:"task"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

func (r *Service) GetPoints(ctx context.Context, runID uint64) ([]Point, error) {
	stats, err := r.repository.GetStats(ctx, runID)
	if err != nil {
		return nil, errors.Wrap(err, "could not get stats")
	}

	var points []Point
	for _, stat := range stats {
		points = append(points, Point{
			RunID: stat.RunID,
			Idx:   stat.Idx,
			Index: r.NameToIndex[stat.Index].Info.LongName,
			Task:  r.NameToTask[stat.Task].Info.LongName,
			X:     float64(stat.Amount),
			Y:     float64(stat.Durs.Median().Microseconds()),
		})
	}

	return points, nil
}
