package cases

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases/stats"
	"geoindexing_comparison/core/cases/tasks"
	"geoindexing_comparison/core/generator"
	"geoindexing_comparison/core/geo"
	"github.com/rs/zerolog/log"
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
	const repetitions = 10

	dur := make([]time.Duration, repetitions)
	for i := 0; i < repetitions; i++ {
		col := colInit()
		col.FromArray(points)

		runtime.GC()
		dur[i] = task.Run(col, amount)
		runtime.GC()
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
	for amount := runCase.AmountStart; amount < runCase.AmountEnd; amount += runCase.AmountStep {
		points := generator.DefaultGenerator.Points(&generator.DefaultInput, amount)

		for _, colInit := range runCase.Collections {
			results = append(results, runCol(points, colInit, amount, task))
		}

		log.Info().
			Str("status", "task.amount.done").
			Str("task", task.Name()).
			Int("amount", amount).
			Send()
	}

	return results
}

func Run(runCase *RunCase) {
	t0 := time.Now()
	for _, task := range runCase.Tasks {
		t1 := time.Now()
		results := runTask(task, runCase)

		log.Info().
			Str("status", "task.done").
			Str("task", task.Name()).
			Dur("elapsed", time.Since(t1)).
			Send()

		drawResultsForTask(task, runCase.Name, results)
	}

	log.Info().
		Str("status", "run.done").
		Dur("elapsed", time.Since(t0)).
		Send()
}
