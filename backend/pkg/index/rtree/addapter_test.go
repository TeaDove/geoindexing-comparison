package rtree

import (
	"geoindexing_comparison/pkg/generator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_RTree_GeneratePoint_Ok(t *testing.T) {
	t.Parallel()

	points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result, _ := collection.KNNTimed(origin, 1000)
	assert.Len(t, result, 1000)
}

func TestUnit_RTree_FindRange_Ok(t *testing.T) {
	t.Parallel()

	points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result, _ := collection.KNNTimed(origin, 1000)
	assert.Len(t, result, 1000)
}

func TestUnit_RTree_RangeSearch_Ok(t *testing.T) {
	t.Parallel()

	points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	collection.BBoxTimed(points.GetRandomPoint(), points.GetRandomPoint())
}
