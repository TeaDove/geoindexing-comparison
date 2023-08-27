package kdtree

import (
	"testing"

	"geoindexing_comparison/generator"
	"geoindexing_comparison/geo"
)

var points = generator.DefaultGenerator.GeneratePointsDefaultAmount()

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	collection := CollectionKDTree{}
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result := collection.KNN(origin, generator.DefaultGenerator.KNNSearchSize)

	collection.Points().
		ToPointExtended().
		PaintPartially(geo.Green, result).
		PaintPartially(geo.Red, []geo.Point{origin}).
		MustExport(geo.ExportInput{})
}

func TestUnit_KDTree_FindRange_Ok(t *testing.T) {
	collection := CollectionKDTree{}
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result := collection.KNN(origin, generator.DefaultGenerator.KNNSearchSize)

	collection.Points().
		ToPointExtended().
		PaintPartially(geo.Green, result).
		PaintPartially(geo.Red, []geo.Point{origin}).
		MustExport(geo.ExportInput{})
}
