package task

import (
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/service/stats"
	"runtime"
	"time"
)

type Insert struct{}

func (r *Insert) Run(input *Input) time.Duration {
	durs := make([]time.Duration, 0)

	for range 20 {
		runtime.GC()

		durs = append(
			durs,
			input.Index.InsertTimed(generator.DefaultGenerator.Point(&generator.DefaultInput)),
		)

		runtime.GC()
	}

	return stats.NewArray(durs).Median()
}
