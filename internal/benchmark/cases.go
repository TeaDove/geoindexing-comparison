package benchmark

import (
	"geoindexing_comparison/generator"
	"geoindexing_comparison/geo"
	"github.com/guregu/null"
)

func (r *Benchmark) KNN() {
	points := r.Generator.GeneratePointsDefaultAmount()
	(*r.Collection).FromArray(points)

	origin := points.GetRandomPoint()
	result := (*r.Collection).KNN(origin, generator.DefaultGenerator.KNNSearchSize)

	(*r.Collection).Points().PaintPartially(geo.Green, result).PaintPartially(geo.Red, []geo.Point{origin}).MustDraw(geo.DrawConfig{OperationType: null.StringFrom("KNN")})
}

func (r *Benchmark) Search() {
	points := r.Generator.GeneratePointsDefaultAmount()
	(*r.Collection).FromArray(points)

	origin := points.GetRandomPoint()
	result := (*r.Collection).RangeSearch(origin, generator.DefaultGenerator.RadiusSearchSize)

	(*r.Collection).Points().PaintPartially(geo.Green, result).PaintPartially(geo.Red, []geo.Point{origin}).MustDraw(geo.DrawConfig{OperationType: null.StringFrom("Search")})
}
