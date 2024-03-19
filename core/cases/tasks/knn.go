package tasks

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/generator"
	"time"
)

type KNN struct{}

func (r *KNN) Name() string {
	return "SimpleKNN"
}

func (r *KNN) Description() string {
	return ""
}

func (r *KNN) Run(collection addapter_all.CollectionInit, amount int) time.Duration {
	col := collection()
	col.FromArray(generator.DefaultGenerator.Points(&generator.DefaultInput, amount))
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	_, t := col.KNNTimed(point, 10)

	return t
}
