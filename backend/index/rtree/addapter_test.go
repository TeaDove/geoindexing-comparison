package rtree

import (
	"geoindexing_comparison/backend/generator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_RTree_GeneratePoint_Ok(t *testing.T) {
	t.Parallel()

	points := generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result, _ := collection.KNNTimed(origin, 1000)
	assert.Len(t, result, 1000)
}

func TestUnit_RTree_FindRange_Ok(t *testing.T) {
	t.Parallel()

	points := generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result, _ := collection.KNNTimed(origin, 1000)
	assert.Len(t, result, 1000)
}

func TestUnit_RTree_RangeSearch_Ok(t *testing.T) {
	t.Parallel()

	points := generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	collection.RangeSearchTimed(
		origin,
		(generator.DefaultInput.LatUpperBound-generator.DefaultInput.LatLowerBound)/6,
	)
}
