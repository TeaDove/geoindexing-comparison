package rtree

import (
	"geoindexing_comparison/core/generator"
	"geoindexing_comparison/core/geo"
	"github.com/stretchr/testify/assert"
	"testing"
)

var points = generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)

func TestUnit_RTree_GeneratePoint_Ok(t *testing.T) {
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result, _ := collection.KNNTimed(origin, 1000)
	assert.Len(t, result, 1000)

	points.
		ToPointExtended().
		PaintPartially(geo.Green, result).
		PaintPartially(geo.Red, []geo.Point{origin})
}

func TestUnit_RTree_FindRange_Ok(t *testing.T) {
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result, _ := collection.KNNTimed(origin, 1000)
	assert.Len(t, result, 1000)

	points.
		ToPointExtended().
		PaintPartially(geo.Green, result).
		PaintPartially(geo.Red, []geo.Point{origin})
}

func TestUnit_RTree_RangeSearch_Ok(t *testing.T) {
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result, _ := collection.RangeSearchTimed(
		origin,
		(generator.DefaultInput.LatUpperBound-generator.DefaultInput.LatLowerBound)/6,
	)
	println(len(result), len(points))
}
