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
	result.Paint(geo.FOUND)
	print(result.String())

	collection.Points().MustExport("points.csv")
}
