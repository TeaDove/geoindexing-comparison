package addapter

import (
	"geoindexing_comparison/generator"
	"testing"
)

var points = generator.DefaultGenerator.GeneratePoints(10)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	collection := CollectionKDTree{}
	collection.FromArray(points)
	print(collection.String())
}
