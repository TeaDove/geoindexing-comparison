package tasks

import (
	"geoindexing_comparison/addapter/addapter_all"
	"geoindexing_comparison/generator"
	"time"
)

type Insert struct{}

func (r *Insert) Name() string {
	return "SimpleInsert"
}

func (r *Insert) Description() string {
	return ""
}

func (r *Insert) Run(collection addapter_all.CollectionInit, amount int) time.Duration {
	col := collection()
	col.FromArray(generator.DefaultGenerator.Points(&generator.DefaultInput, amount))
	point := generator.DefaultGenerator.Point(&generator.DefaultInput)

	t := col.InsertTimed(point)

	return t
}
