package kdtree

import (
	"geoindexing_comparison/service/generator"
	"testing"
)

var points = generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	t.Parallel()

	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	collection.KNNTimed(origin, 10)
}
