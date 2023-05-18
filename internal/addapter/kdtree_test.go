package addapter

import (
	"geoindexing_comparison/generator"
	"geoindexing_comparison/geo"
	"testing"
)

var points = generator.DefaultGenerator.GeneratePointsDefaultAmount()

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	collection := CollectionKDTree{}
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result := collection.KNN(origin, generator.DefaultGenerator.KNNSearchSize)

	collection.Points().PaintPartially(geo.Green, result).PaintPartially(geo.Red, []geo.Point{origin}).MustExport("points.csv")
}

func TestUnit_KDTree_FindRange_Ok(t *testing.T) {
	collection := CollectionKDTree{}
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result := collection.RangeSearch(origin, generator.DefaultGenerator.RadiusSearchSize)

	collection.Points().PaintPartially(geo.Green, result).PaintPartially(geo.Red, []geo.Point{origin}).Draw()
}
