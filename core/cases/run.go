package cases

import (
	"geoindexing_comparison/core/cases/stats"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

const repetitions = 30

func runAmount(wg *sync.WaitGroup, runCase RunCase, amount int) {
	dur := make([]time.Duration, repetitions)
	for i := 0; i < repetitions; i++ {
		dur[i] = runCase.Task.Run(runCase.Collection, amount)

	}
	result := stats.NewDurs(dur)

	log.Info().
		Str("status", "run.done").
		Str("collection", runCase.Collection().Name()).
		Str("task", runCase.Task.Name()).
		Str("result", result.String()).
		Int("amount", amount).
		Send()

	wg.Done()
}

func run(wg *sync.WaitGroup, resultChan chan stats.Durs, runCase RunCase) {
	defer wg.Done()

	log.Info().
		Str("status", "run.started").
		Str("collection", runCase.Collection().Name()).
		Str("task", runCase.Task.Name()).
		Send()

	cur := runCase.AmountStart
	for {
		if cur > runCase.AmountEnd {
			break
		}

		wg.Add(1)
		go runAmount(wg, runCase, cur)

		cur += runCase.AmountStep
	}

	//resultChan <- result
}

func Run(runCases ...RunCase) {
	var wg sync.WaitGroup
	resultChan := make(chan stats.Durs, len(runCases))

	for _, runCase := range runCases {
		wg.Add(1)
		go run(&wg, resultChan, runCase)
	}

	wg.Wait()
	close(resultChan)
}
