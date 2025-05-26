package task

import (
	"geoindexing_comparison/pkg/generator"
	"geoindexing_comparison/pkg/helpers/stats"
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
			input.Index.InsertTimed(generator.NewSimplerGenerator().Point(&generator.DefaultInput)),
		)

		runtime.GC()
	}

	return stats.NewArray(durs).Median()
}
