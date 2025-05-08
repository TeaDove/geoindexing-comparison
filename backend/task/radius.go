package task

import (
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/index"
	"time"
)

type BBox struct{}

func (r *BBox) Run(index index.Impl, _ uint64) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := index.BBoxTimed(
		point,
		generator.DefaultGenerator.Point(&generator.DefaultInput),
	)

	return t
}
