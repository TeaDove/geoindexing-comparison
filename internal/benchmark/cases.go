package benchmark

import (
	"geoindexing_comparison/generator"
	"geoindexing_comparison/geo"

	"github.com/guregu/null"
)

func (r *Benchmark) KNNCheck() {
	points := r.Generator.GeneratePointsDefaultAmount()
	r.Collection.FromArray(points)

	origin := points.GetRandomPoint()
	result := r.Collection.KNN(origin, generator.DefaultGenerator.KNNSearchSize)

	r.Collection.Points().
		ToPointExtended().
		PaintPartially(geo.Green, result).
		PaintPartially(geo.Red, []geo.Point{origin}).
		MustDraw(geo.DrawInput{OperationType: null.StringFrom("KNN")})
}

func (r *Benchmark) SearchCheck() {
	points := r.Generator.GeneratePointsDefaultAmount()
	r.Collection.FromArray(points)

	origin := points.GetRandomPoint()
	result := r.Collection.RangeSearch(origin, generator.DefaultGenerator.RadiusSearchSize)

	r.Collection.Points().
		ToPointExtended().
		PaintPartially(geo.Green, result).
		PaintPartially(geo.Red, []geo.Point{origin}).
		MustDraw(geo.DrawInput{OperationType: null.StringFrom("Search")})
}
