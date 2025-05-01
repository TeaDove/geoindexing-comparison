package quadtree

import (
	"geoindexing_comparison/backend/generator"
	"testing"
)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	t.Parallel()

	points := generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	collection.KNNTimed(origin, 10)
}
