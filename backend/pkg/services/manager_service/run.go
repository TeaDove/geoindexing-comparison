package manager_service

import (
	"context"
	"geoindexing_comparison/pkg/generator"
	"geoindexing_comparison/pkg/repositories/manager_repository"
	"geoindexing_comparison/pkg/schemas"
	"github.com/guregu/null/v6"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func (r *Service) generateJobs(run *manager_repository.Run) {
	r.jobMu.Lock()
	defer r.jobMu.Unlock()

	r.jobs = make([]schemas.Job, 0)

	for amount := run.Start; amount < run.Stop; amount += run.Step {
		for _, runTask := range run.Tasks {
			points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, amount)
			point := points.GetRandomPoint()

			for _, runIndex := range run.Indexes {
				r.jobs = append(r.jobs, schemas.Job{
					Task:        runTask,
					Index:       runIndex,
					Amount:      amount,
					Points:      points,
					RandomPoint: point,
				})
			}
		}
	}

	r.jobIdx = 0
	r.currentRun = run
}

func (r *Service) runPending(ctx context.Context, run *manager_repository.Run) error {
	ctx = logger_utils.WithValue(ctx, "run_id", strconv.Itoa(run.ID))

	if len(run.Tasks) == 0 || len(run.Indexes) == 0 {
		run.CompletedAt = null.TimeFrom(time.Now())
		run.Status = manager_repository.RunStatusCancelled
		err := r.repository.SaveRun(ctx, run)
		if err != nil {
			return errors.Wrap(err, "failed to save run")
		}

		zerolog.Ctx(ctx).Warn().Msg("empty.run")
		return nil
	}

	zerolog.Ctx(ctx).Info().Interface("run", run).Msg("run.started")

	r.generateJobs(run)

	for {
		curRun, err := r.repository.GetRun(ctx, run.ID)
		if err != nil {
			return errors.Wrap(err, "failed to get current run")
		}

		if curRun.Status != manager_repository.RunStatusPending {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func (r *Service) initRunner() {
	const sleepOnErr = time.Second * 5

	for {
		ctx := logger_utils.NewLoggedCtx()

		pendingRuns, err := r.repository.GetPending(ctx)
		if err != nil {
			zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("failed.to.get.pending.runs")
			time.Sleep(sleepOnErr)

			continue
		}

		for _, run := range pendingRuns {
			err = r.runPending(ctx, &run)
			if err != nil {
				zerolog.Ctx(ctx).Error().
					Stack().
					Err(err).
					Interface("run", run).
					Msg("failed.to.run")
				time.Sleep(sleepOnErr)
			}
		}

		time.Sleep(time.Second)
	}
}
