package manager_service

import (
	"context"
	"geoindexing_comparison/pkg/repositories/manager_repository"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type RunRequest struct {
	Indexes []string `json:"indexes"`
	Tasks   []string `json:"tasks"`
	Start   int      `json:"start"`
	Stop    int      `json:"stop"`
	Step    int      `json:"step"`
}

func (r *Service) AddRun(ctx context.Context, req *RunRequest, createdBy string) (*manager_repository.Run, error) {
	run := &manager_repository.Run{
		CreatedAt: time.Now().UTC(),
		CreatedBy: createdBy,
		Status:    manager_repository.RunStatusPending,
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
	r.jobMu.Lock()
	defer r.jobMu.Unlock()

	r.currentRun = nil
	r.jobIdx = 0
	r.jobs = nil

	err := r.repository.StopRuns(ctx)
	if err != nil {
		return errors.Wrap(err, "could not save run")
	}

	zerolog.Ctx(ctx).Info().Msg("runs.stopped")

	return nil
}

func (r *Service) GetRuns(ctx context.Context) ([]manager_repository.Run, error) {
	return r.repository.GetRuns(ctx)
}

type Point struct {
	RunID int `json:"runId"`
	Idx   int `json:"idx"`

	Index string  `json:"index"`
	Task  string  `json:"task"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

func (r *Service) GetChartPoints(ctx context.Context, runID int) ([]Point, error) {
	stats, err := r.repository.GetStats(ctx, runID)
	if err != nil {
		return nil, errors.Wrap(err, "could not get stats")
	}

	var points []Point
	for _, stat := range stats {
		points = append(points, Point{
			RunID: stat.RunID,
			Idx:   stat.Idx,
			Index: r.builderService.IndexMap[stat.Index].Info.LongName,
			Task:  r.builderService.TaskMap[stat.Task].Info.LongName,
			X:     float64(stat.Amount),
			Y:     stat.Durs.QualifiedAvg() / float64(time.Microsecond),
		})
	}

	return points, nil
}
