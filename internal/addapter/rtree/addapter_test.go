package kdtree

import (
	"geoindexing_comparison/generator"
	"geoindexing_comparison/geo"
	"testing"
)

var points = generator.DefaultGenerator.GeneratePointsDefaultAmount()

func TestUnit_RTree_GeneratePoint_Ok(t *testing.T) {
	collection := CollectionRTree{}
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result := collection.KNN(origin, generator.DefaultGenerator.KNNSearchSize)

	collection.Points().
		ToPointExtended().
		PaintPartially(geo.Green, result).
		PaintPartially(geo.Red, []geo.Point{origin}).
		MustExport(geo.ExportInput{})
}

func TestUnit_RTree_FindRange_Ok(t *testing.T) {
	collection := CollectionRTree{}
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result := collection.KNN(origin, generator.DefaultGenerator.KNNSearchSize)

	collection.Points().
		ToPointExtended().
		PaintPartially(geo.Green, result).
		PaintPartially(geo.Red, []geo.Point{origin}).
		MustExport(geo.ExportInput{})
}
