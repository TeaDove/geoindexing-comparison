package cases

import (
	"geoindexing_comparison/service/cases/stats"
	"geoindexing_comparison/service/cases/tasks"
	"geoindexing_comparison/service/generator"
	"geoindexing_comparison/service/geo"
	"geoindexing_comparison/service/index"
	"geoindexing_comparison/service/index/indexes"
	"runtime"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
)

type Result struct {
	IndexName string
	Task      tasks.TaskImpl
	Durs      stats.Array[time.Duration]
	Mems      stats.Array[uint64]
	Amount    uint64
}

func runCol(
	points geo.Points,
	colInit index.NewIndex,
	amount uint64,
	task tasks.TaskImpl,
) Result {
	const repetitions = 5

	dur := make([]time.Duration, repetitions)
	mems := make([]uint64, repetitions)

	for i := 0; i < repetitions; i++ {
		col := colInit()
		col.FromArray(points)

		runtime.GC()

		// TODO add mem checker

		dur[i] = task.Run(col, amount)

		runtime.GC()

		if dur[i].Seconds() > 0.5 {
			log.Warn().
				Str("status", "run.is.too.long").
				Dur("dur", dur[i]).
				Str("col", col.Name()).
				Str("task", task.Filename()).
				Send()

			break
		}
	}

	return Result{
		IndexName: colInit().Name(),
		Task:      task,
		Durs:      stats.NewArray(dur),
		Amount:    amount,
		Mems:      stats.NewArray(mems),
	}
}

func runTask(task tasks.TaskImpl, runCase *RunCase) []Result {
	results := make([]Result, 0, 10)
	iterations := (runCase.AmountEnd - runCase.AmountStart) / runCase.AmountStep
	bar := progressbar.Default(int64(iterations))

	for amount := runCase.AmountStart; amount < runCase.AmountEnd; amount += runCase.AmountStep {
		// points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)
		points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)

		for _, colInit := range runCase.Indexes {
			results = append(results, runCol(points, indexes.NameToNewIndex[colInit], amount, task))
		}

		err := bar.Add(1)
		if err != nil {
			log.Error().Err(err).Str("status", "bar.add.error")
		}
	}

	return results
}

func Run(runCase *RunCase) {
	t0 := time.Now()

	for _, task := range runCase.Tasks {
		t1 := time.Now()

		log.Info().
			Str("status", "task.begin").
			Str("task", task).
			Send()

		_ = runTask(tasks.NameToTask[task], runCase)

		log.Info().
			Str("status", "task.done").
			Str("task", task).
			Str("elapsed", time.Since(t1).String()).
			Send()
	}

	log.Info().
		Str("status", "run.done").
		Str("elapsed", time.Since(t0).String()).
		Send()
}
