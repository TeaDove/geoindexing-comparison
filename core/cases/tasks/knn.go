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

type KNN90 struct{}

func (r *KNN90) Name() string {
	return "КНН"
}

func (r *KNN90) Filename() string { return "knn_90" }

func (r *KNN90) Description() string {
	return "КНН на 90% точек из структуры"
}

func (r *KNN90) Run(col addapter.Collection, amount int) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.KNNTimed(point, amount*90/100)

	return t
}

type KNN1 struct{}

func (r *KNN1) Name() string {
	return "КНН"
}

func (r *KNN1) Filename() string { return "knn_1" }

func (r *KNN1) Description() string {
	return "КНН на 1% точек из структуры"
}

func (r *KNN1) Run(col addapter.Collection, amount int) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.KNNTimed(point, amount/100)

	return t
}
