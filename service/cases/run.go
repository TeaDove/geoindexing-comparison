package cases

import (
	"context"
	"geoindexing_comparison/service/cases/stats"
	"geoindexing_comparison/service/cases/tasks"
	"geoindexing_comparison/service/generator"
	"geoindexing_comparison/service/geo"
	"geoindexing_comparison/service/index"
	"github.com/rs/zerolog"
	"runtime"
	"time"
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

	dur := make([]time.Duration, repetitions)
	mems := make([]uint64, repetitions)

	for i := 0; i < repetitions; i++ {
		idxImpl := idx.Builder()
		idxImpl.FromArray(points)

		runtime.GC()

		// TODO add mem checker

		dur[i] = task.Builder().Run(idxImpl, amount)

		runtime.GC()

		if dur[i].Seconds() > 0.5 {
			zerolog.Ctx(ctx).
				Warn().
				Dur("dur", dur[i]).
				Str("col", idx.Info.ShortName).
				Str("task", task.Info.ShortName).
				Msg("run.is.too.long")

			break
		}
	}

	return Result{
		Index:  idx.Info.ShortName,
		Task:   task.Info.ShortName,
		Durs:   stats.NewArray(dur),
		Amount: amount,
		Mems:   stats.NewArray(mems),
	}
}

func (r *Runner) runTask(ctx context.Context, task tasks.Task, runCase *RunConfig, channel chan<- Result) {
	for amount := runCase.AmountStart; amount < runCase.AmountEnd; amount += runCase.AmountStep {
		if ctx.Err() != nil {
			return
		}
		// points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)
		points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)

		for _, colInit := range runCase.Indexes {
			channel <- runCol(ctx, points, r.NameToIndex[colInit], task, amount)
		}
	}
}

func (r *Runner) Run(ctx context.Context, runCase *RunConfig) <-chan Result {
	ctx, cancel := context.WithCancel(ctx)

	t0 := time.Now()
	channel := make(chan Result, 1000)
	go func() {
		for _, task := range runCase.Tasks {
			if ctx.Err() != nil {
				close(channel)
				return
			}

			t1 := time.Now()

			zerolog.Ctx(ctx).Info().
				Str("task", task).
				Msg("task.begin")

			r.runTask(ctx, r.NameToTask[task], runCase, channel)

			zerolog.Ctx(ctx).Info().
				Str("task", task).
				Str("elapsed", time.Since(t1).String()).
				Msg("task.done")
		}

		zerolog.Ctx(ctx).Info().
			Str("elapsed", time.Since(t0).String()).
			Msg("run.done")
	}()

	r.stop = cancel

	return channel
}

func (r *Runner) Stop(ctx context.Context) {
	if r.stop == nil {
		zerolog.Ctx(ctx).Warn().Msg("runner.is.already.stopped")
	}
	r.stop()
}
