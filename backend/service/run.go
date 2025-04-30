package service

import (
	"context"
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"geoindexing_comparison/backend/repository"
	"geoindexing_comparison/backend/service/stats"
	"geoindexing_comparison/backend/tasks"
	"runtime"
	"strconv"
	"time"

	"github.com/guregu/null/v6"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/logger_utils"
	"github.com/teadove/teasutils/utils/time_utils"
)

type Result struct {
	Index  string
	Task   string
	Amount uint64
	Durs   stats.Array[time.Duration]
	Mems   stats.Array[uint64]
}

func runCol(ctx context.Context,
	points geo.Points,
	idx index.Index,
	task tasks.Task,
	amount uint64,
) Result {
	const repetitions = 5

	durs := make([]time.Duration, 0)
	mems := make([]uint64, 0)

	for range repetitions {
		idxImpl := idx.Builder()
		idxImpl.FromArray(points)

		runtime.GC()

		// TODO add mem checker

		dur := task.Builder().Run(idxImpl, amount)

		runtime.GC()

		durs = append(durs, dur)

		if dur > 1*time.Second {
			zerolog.Ctx(ctx).
				Warn().
				Str("dur", dur.String()).
				Str("col", idx.Info.ShortName).
				Str("task", task.Info.ShortName).
				Msg("run.is.too.long")

			break
		}
	}

	return Result{
		Index:  idx.Info.ShortName,
		Task:   task.Info.ShortName,
		Durs:   stats.NewArray(durs),
		Amount: amount,
		Mems:   stats.NewArray(mems),
	}
}

func (r *Service) run(ctx context.Context, run *repository.Run) error {
	var idx uint64

	for _, task := range run.Tasks {
		t1 := time.Now()

		zerolog.Ctx(ctx).Debug().
			Str("task", task).
			Msg("task.begin")

		for amount := run.Start; amount < run.Stop; amount += run.Step {
			// points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)
			points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)

			for _, runIndex := range run.Indexes {
				result := runCol(ctx, points, r.NameToIndex[runIndex], r.NameToTask[task], amount)

				err := r.repository.SaveStats(ctx, &repository.Stats{
					Idx:    idx,
					RunID:  run.ID,
					Index:  result.Index,
					Task:   result.Task,
					Amount: result.Amount,
					Durs:   result.Durs,
				})
				if err != nil {
					return errors.Wrap(err, "failed to save")
				}

				if idx%10 == 0 {
					zerolog.Ctx(ctx).Debug().
						Uint64("idx", idx).
						Str("task", task).
						Msg("iteration.done")
				}

				idx++
			}
		}

		zerolog.Ctx(ctx).Debug().
			Str("task", task).
			Str("elapsed", time_utils.RoundDuration(time.Since(t1))).
			Msg("task.done")
	}

	return nil
}

const sleepOnErr = time.Second * 5

func (r *Service) initRunner() {
	for {
		ctx := logger_utils.NewLoggedCtx()

		pendingRuns, err := r.repository.GetPending(ctx)
		if err != nil {
			zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("failed.to.get.pending.runs")
			time.Sleep(sleepOnErr)

			continue
		}

		for _, run := range pendingRuns {
			ctx := logger_utils.WithValue(ctx, "run_id", strconv.FormatUint(run.ID, 10))
			zerolog.Ctx(ctx).Info().Interface("run", run).Msg("run.started")

			t0 := time.Now()

			err = r.run(ctx, &run)
			if err != nil {
				zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("failed.to.run")
				time.Sleep(sleepOnErr)

				continue
			}

			run.CompletedAt = null.NewTime(time.Now(), true)
			run.Status = repository.RunStatusCompleted

			err = r.repository.SaveRun(ctx, &run)
			if err != nil {
				zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("failed.to.save.run")
				time.Sleep(sleepOnErr)

				continue
			}

			zerolog.Ctx(ctx).Info().
				Str("elapsed", time_utils.RoundDuration(time.Since(t0))).
				Msg("run.done")
		}

		time.Sleep(time.Second)
	}
}
