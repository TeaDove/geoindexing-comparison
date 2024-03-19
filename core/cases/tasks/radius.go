package tasks

import (
	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/generator"
	"time"
)

type RadiusSearch struct {
	collection addapter.Collection
}

func (r *RadiusSearch) Name() string {
	return "RadiusSearch"
}

func (r *RadiusSearch) Description() string {
	return "DefaultGenerator"
}

func (r *RadiusSearch) Run(collection addapter_all.CollectionInit, amount int) time.Duration {
	col := collection()

	col.FromArray(generator.DefaultGenerator.Points(&generator.DefaultInput, amount))
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.RangeSearchTimed(point, 1000)

	return t
}
