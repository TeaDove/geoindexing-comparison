package cases

import (
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

func run(wg *sync.WaitGroup, resultChan chan Result, runCase RunCase) {
	defer wg.Done()

	log.Info().
		Str("status", "run.started").
		Str("collection", runCase.Collection().Name()).
		Str("task", runCase.Task.Name()).
		Send()

	dur := make([]time.Duration, runCase.Repetitions)

	for i := 0; i < runCase.Repetitions; i++ {
		dur[i] = runCase.Task.Run(runCase.Collection)
	}
	result := NewResult(dur)

	log.Info().
		Str("status", "run.done").
		Str("collection", runCase.Collection().Name()).
		Str("task", runCase.Task.Name()).
		Str("result", result.String()).
		Send()

	resultChan <- result
}

func Run(runCases ...RunCase) {
	var wg sync.WaitGroup
	resultChan := make(chan Result, len(runCases))

	for _, runCase := range runCases {
		wg.Add(1)
		go run(&wg, resultChan, runCase)
	}

	wg.Wait()
	close(resultChan)
}
