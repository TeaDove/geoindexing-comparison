package tasks

import (
	"geoindexing_comparison/service/cases/stats"
	"geoindexing_comparison/service/generator"
	"geoindexing_comparison/service/index"
	"runtime"
	"time"
)

type Insert struct{}

func (r *Insert) Name() string {
	return "Вставка"
}

func (r *Insert) Filename() string { return "insert" }

func (r *Insert) Description() string {
	return "Вставка 10% точек"
}

func (r *Insert) Run(index index.IndexImpl, amount uint64) time.Duration {
	durs := make([]time.Duration, 0, amount/10)

	for range amount / 10 {
		runtime.GC()

		durs = append(
			durs,
			index.InsertTimed(generator.DefaultGenerator.Point(&generator.DefaultInput)),
		)

		runtime.GC()
	}

	return stats.NewArray(durs).Avg()
}
