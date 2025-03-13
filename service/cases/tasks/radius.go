package tasks

import (
	"geoindexing_comparison/service/generator"
	"geoindexing_comparison/service/index"
	"time"
)

type RadiusSearch struct{}

func (r *RadiusSearch) Name() string {
	return "Поиск в радиусе"
}

func (r *RadiusSearch) Filename() string { return "radius_search" }

func (r *RadiusSearch) Description() string {
	return ""
}

func (r *RadiusSearch) Run(index index.Index, amount uint64) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := index.RangeSearchTimed(
		point,
		(generator.DefaultInput.LatUpperBound-generator.DefaultInput.LatLowerBound)/6,
	)

	return t
}
