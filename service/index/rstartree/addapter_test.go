package rstartree

import (
	"geoindexing_comparison/service/generator"
	"geoindexing_comparison/service/geo"
	"testing"

	"github.com/stretchr/testify/assert"
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
	t.Parallel()

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
