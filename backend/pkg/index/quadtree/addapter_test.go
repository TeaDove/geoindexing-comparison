package quadtree

import (
	"geoindexing_comparison/pkg/generator"
	"testing"
)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	t.Parallel()

	points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	collection.KNNTimed(origin, 10)
}
