package tasks

import (
	"geoindexing_comparison/addapter"
	"geoindexing_comparison/generator"
	"time"
)

type RadiusSearch struct {
	collection addapter.Collection
}

func (r *RadiusSearch) Name() string {
	return "RadiusSearch"
}

func (r *RadiusSearch) Run(collection func() addapter.Collection) time.Duration {
	col := collection()
	col.FromArray(generator.DefaultGenerator.GeneratePoints(100_000))
	point := generator.DefaultGenerator.GeneratePoint()

	t0 := time.Now()
	_ = col.RangeSearch(point, 10)

	return time.Now().Sub(t0)
}
