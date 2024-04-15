package tasks

import (
	"runtime"
	"time"

	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/cases/stats"
	"geoindexing_comparison/core/generator"
)

type Insert struct{}

func (r *Insert) Name() string {
	return "Вставка"
}

func (r *Insert) Filename() string { return "insert" }

func (r *Insert) Description() string {
	return "Вставка 10% точек"
}

func (r *Insert) Run(col addapter.Collection, amount int) time.Duration {
	durs := make([]time.Duration, 0, amount/10)

	for range amount / 10 {
		runtime.GC()
		durs = append(
			durs,
			col.InsertTimed(generator.DefaultGenerator.Point(&generator.DefaultInput)),
		)
		runtime.GC()
	}

	return stats.NewDurs(durs).Avg()
}
