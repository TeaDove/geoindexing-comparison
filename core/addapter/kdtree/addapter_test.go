package kdtree

import (
	"testing"

	"geoindexing_comparison/core/generator"
)

var points = generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	collection.KNNTimed(origin, 10)
}
