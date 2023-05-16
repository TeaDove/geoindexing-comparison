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

	result := collection.KNN(generator.DefaultGenerator.GeneratePoint(), 10)

	collection.Points().PaintPartially(geo.FOUND, result).MustExport("points.csv")
}
