package cases

import (
	"geoindexing_comparison/core/cases/stats"
	"github.com/rs/zerolog/log"
	"runtime"
	"time"
)

const repetitions = 30

func runAmount(runCase RunCase, amount int) Result {
	dur := make([]time.Duration, repetitions)
	for i := 0; i < repetitions; i++ {
		dur[i] = runCase.Task.Run(runCase.Collection, amount)

	}
	result := stats.NewDurs(dur)

	runtime.GC()

	return Result{
		RunCase: runCase,
		Durs:    result,
		Amount:  amount,
	}
}

type Result struct {
	RunCase RunCase
	Durs    stats.Durs
	Amount  int
}

func Run(runCases []RunCase) []Result {
	results := make([]Result, 0, 100)
	for _, runCase := range runCases {
		cur := runCase.AmountStart
		for {
			if cur > runCase.AmountEnd {
				break
			}

			results = append(results, runAmount(runCase, cur))

			cur += runCase.AmountStep
		}

		log.Info().
			Str("status", "сase.done").
			Str("collection", runCase.Collection().Name()).
			Str("task", runCase.Task.Name()).
			Send()
	}

	return results
}
