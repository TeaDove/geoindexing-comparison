package worker_service

import (
	"context"
	"geoindexing_comparison/backend/helpers/stats"
	"geoindexing_comparison/backend/schemas"
	"geoindexing_comparison/backend/suppliers/manager_supplier"
	"geoindexing_comparison/backend/task"
	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

func (r *Service) run(ctx context.Context, job *schemas.Job) schemas.JobResult {
	taskObj, ok := r.builderService.NameToTask[job.Task]
	if !ok {
		panic(errors.New("task not found"))
	}

	indexObj, ok := r.builderService.NameToIndex[job.Index]
	if !ok {
		panic(errors.New("index not found"))
	}

	const repetitions = 5

	durs := make([]time.Duration, 0)

	for range repetitions {
		indexImpl := indexObj.Builder()
		indexImpl.FromArray(job.Points)

		taskInput := task.Input{
			Index:       indexImpl,
			Amount:      job.Amount,
			Points:      job.Points,
			RandomPoint: job.RandomPoint,
		}
		taskImpl := taskObj.Builder()

		runtime.GC()

		dur := taskImpl.Run(&taskInput)

		durs = append(durs, dur)

		if dur > 1*time.Second {
			zerolog.Ctx(ctx).
				Warn().
				Str("dur", dur.String()).
				Str("index", indexObj.Info.ShortName).
				Str("task", taskObj.Info.ShortName).
				Int("amount", job.Amount).
				Msg("run.is.too.long")

			break
		}
	}

	return schemas.JobResult{
		Index:  job.Index,
		Task:   job.Task,
		Amount: job.Amount,
		Durs:   stats.NewArray(durs),
	}
}

func (r *Service) Job() {
	const (
		notFoundSleep = 200 * time.Millisecond
		errorSleep    = 500 * time.Millisecond
	)

	for {
		ctx := logger_utils.NewLoggedCtx()
		job, err := r.managerSupplier.GetPendingJobs(ctx)
		if err != nil {
			if errors.Is(err, manager_supplier.NotFoundError) {
				zerolog.Ctx(ctx).
					Debug().
					Msg("no.pending.jobs")
				time.Sleep(notFoundSleep)
				continue
			}

			zerolog.Ctx(ctx).
				Error().
				Stack().Err(err).
				Msg("failed.to.get.pending.jobs")
			time.Sleep(errorSleep)
			continue
		}

		res := r.run(ctx, &job)

		err = r.managerSupplier.ReportJob(ctx, res)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Stack().Err(err).
				Msg("failed.to.report.result")
			time.Sleep(errorSleep)
			continue
		}
	}
}
