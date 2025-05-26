package manager_service

import (
	"context"
	"geoindexing_comparison/pkg/generator"
	"geoindexing_comparison/pkg/repositories/manager_repository"
	"geoindexing_comparison/pkg/schemas"
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

	r.allJobsDone = make(chan struct{})
	r.jobIdx = 0
	r.currentRun = run
}

func (r *Service) runPending(ctx context.Context, run *manager_repository.Run) error {
	ctx = logger_utils.WithValue(ctx, "run_id", strconv.Itoa(run.ID))

	zerolog.Ctx(ctx).Info().Interface("run", run).Msg("run.started")

	r.generateJobs(run)

	<-r.allJobsDone

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
