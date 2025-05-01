package tasks

import (
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/index"
	"time"
)

type KNN25P struct{}

func (r *KNN25P) Run(index index.IndexImpl, amount uint64) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := index.KNNTimed(point, amount/4)

	return t
}

type KNN90P struct{}

func (r *KNN90P) Run(col index.IndexImpl, amount uint64) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.KNNTimed(point, amount*90/100)

	return t
}

type KNN1P struct{}

func (r *KNN1P) Run(col index.IndexImpl, amount uint64) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.KNNTimed(point, amount/100)

	return t
}

type KNN10 struct{}

func (r *KNN10) Run(col index.IndexImpl, _ uint64) time.Duration {
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.KNNTimed(point, 10)

	return t
}
