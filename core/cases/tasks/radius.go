package tasks

import (
	"time"

	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/generator"
)

type RadiusSearch struct{}

func (r *RadiusSearch) Name() string {
	return "Поиск в радиусе"
}

func (r *RadiusSearch) Filename() string { return "radius_search" }

func (r *RadiusSearch) Description() string {
	return ""
}

func (r *RadiusSearch) Run(col addapter.Collection, amount int) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.RangeSearchTimed(
		point,
		(generator.DefaultInput.LatUpperBound-generator.DefaultInput.LatLowerBound)/6,
	)

	return t
}
