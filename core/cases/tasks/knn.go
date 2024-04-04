package tasks

import (
	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/generator"
	"time"
)

type KNNQuater struct{}

func (r *KNNQuater) Name() string {
	return "КНН"
}

func (r *KNNQuater) Filename() string { return "knn_quarters" }

func (r *KNNQuater) Description() string {
	return "КНН на четверть точек"
}

func (r *KNNQuater) Run(col addapter.Collection, amount int) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.KNNTimed(point, amount/4)

	return t
}

type KNNAmountPoint struct{}

func (r *KNNAmountPoint) Name() string {
	return "КНН"
}

func (r *KNNAmountPoint) Filename() string { return "knn_full" }

func (r *KNNAmountPoint) Description() string {
	return "КНН на все точки в структуре"
}

func (r *KNNAmountPoint) Run(col addapter.Collection, amount int) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.KNNTimed(point, amount)

	return t
}
