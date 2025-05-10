package service

import (
	"context"
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"geoindexing_comparison/backend/repository"
	"geoindexing_comparison/backend/service/stats"
	"geoindexing_comparison/backend/task"
	"github.com/teadove/teasutils/utils/must_utils"
	"runtime"
	"strconv"
	"time"

	"gorm.io/gorm"

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
	task task.Task,
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
				Uint64("amount", amount).
				Int("points", len(points)).
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

type RunInput struct {
	Task      task.Task
	Index     index.Index
	TaskInput *task.Input

	Amount uint64
	Points geo.Points
	Point  geo.Point
}

func (r *Service) run(ctx context.Context, run *repository.Run) error {
	var inputs []RunInput

	for amount := run.Start; amount < run.Stop; amount += run.Step {
		for _, runTask := range run.Tasks {
			points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)
			point := points.GetRandomPoint()

			for _, runIndex := range run.Indexes {
				inputs = append(inputs, RunInput{
					Task:   r.NameToTask[runTask],
					Index:  r.NameToIndex[runIndex],
					Amount: amount,
					Points: points,
					Point:  point,
				})
			}
		}
	}

	idx, err := r.repository.GetLastStat(ctx, run.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "failed to get last stat")
	}

	zerolog.Ctx(ctx).
		Info().
		Int("len", len(inputs)).
		Uint64("skipped", idx).
		Msg("inputs.compiled")

	for int(idx) < len(inputs) {
		result := runCol(ctx, inputs[idx].Points, inputs[idx].Index, inputs[idx].Task, inputs[idx].Amount)

		err = r.repository.SaveStats(ctx, &repository.Stats{
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
				Msg("iteration.done")

			run, err = r.repository.GetRun(ctx, run.ID)
			if err != nil {
				return errors.Wrap(err, "failed to get run")
			}
			if run.Status != repository.RunStatusPending {
				zerolog.Ctx(ctx).
					Info().
					Uint64("idx", idx).
					Interface("run", run).
					Msg("ending.run.earlier")
				return nil
			}

		}

		idx++
	}

	return nil
}

func (r *Service) runPending(ctx context.Context, run *repository.Run) error {
	ctx = logger_utils.WithValue(ctx, "run_id", strconv.FormatUint(run.ID, 10))
	defer func() {
		err := must_utils.AnyToErr(recover())
		if err != nil {
			run.Status = repository.RunStatusCancelled
			saveErr := r.repository.SaveRun(ctx, run)
			if saveErr != nil {
				panic("failed to save run")
			}

			zerolog.Ctx(ctx).Error().
				Err(err).
				Interface("run", run).
				Msg("run panicked")
		}

	}()

	zerolog.Ctx(ctx).Info().Interface("run", run).Msg("run.started")

	t0 := time.Now()

	err := r.run(ctx, run)
	if err != nil {
		return errors.Wrap(err, "failed to run")
	}

	run.CompletedAt = null.NewTime(time.Now(), true)
	run.Status = repository.RunStatusCompleted

	err = r.repository.SaveRun(ctx, run)
	if err != nil {
		return errors.Wrap(err, "failed to save")
	}

	zerolog.Ctx(ctx).Info().
		Str("elapsed", time_utils.RoundDuration(time.Since(t0))).
		Msg("run.done")

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
