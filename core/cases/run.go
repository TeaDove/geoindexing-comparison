package cases

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases/stats"
	"geoindexing_comparison/core/cases/tasks"
	"geoindexing_comparison/core/generator"
	"geoindexing_comparison/core/geo"
	"geoindexing_comparison/core/utils"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"runtime"
	"time"
)

type Result struct {
	CollectionName string
	Task           tasks.Task
	Durs           stats.Durs
	Amount         int
}

func runCol(
	points geo.Points,
	colInit addapter_all.CollectionInit,
	amount int,
	task tasks.Task,
) Result {
	const repetitions = 5

	dur := make([]time.Duration, repetitions)
	for i := 0; i < repetitions; i++ {
		col := colInit()
		col.FromArray(points)

		runtime.GC()
		dur[i] = task.Run(col, amount)
		runtime.GC()
		if dur[i].Seconds() > 0.5 {
			log.Warn().
				Str("status", "run.too.long").
				Dur("dur", dur[i]).
				Str("col", col.Name()).
				Str("task", task.Filename()).
				Send()
			break
		}
	}

	return Result{
		CollectionName: colInit().Name(),
		Task:           task,
		Durs:           stats.NewDurs(dur),
		Amount:         amount,
	}
}

func runTask(task tasks.Task, runCase *RunCase) []Result {
	results := make([]Result, 0, 10)
	iterations := (runCase.AmountEnd - runCase.AmountStart) / runCase.AmountStep
	bar := progressbar.Default(int64(iterations))
	for amount := runCase.AmountStart; amount < runCase.AmountEnd; amount += runCase.AmountStep {
		//points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)
		points := generator.DefaultNormalGenerator.Points(&generator.DefaultInput, amount)

		for _, colInit := range runCase.Collections {
			results = append(results, runCol(points, colInit, amount, task))
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
			Str("task", task.Filename()).
			Send()

		results := runTask(task, runCase)

		log.Info().
			Str("status", "task.done").
			Str("task", task.Filename()).
			Float64("elapsed.m", utils.ToFixed(time.Since(t1).Minutes(), 2)).
			Send()

		drawResultsForTask(task, runCase.Name, results)
	}

	log.Info().
		Str("status", "run.done").
		Float64("elapsed.m", utils.ToFixed(time.Since(t0).Minutes(), 2)).
		Send()
}
