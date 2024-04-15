package bruteforce

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"geoindexing_comparison/core/generator"
)

var points = generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	items, _ := collection.KNNTimed(origin, 900)
	assert.Equal(t, 900, len(items))
}

func TestUnit_PointGenerator_RadiusSearch_Ok(t *testing.T) {
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	collection.RangeSearchTimed(
		origin,
		(generator.DefaultInput.LatUpperBound-generator.DefaultInput.LatLowerBound)/6,
	)
}
