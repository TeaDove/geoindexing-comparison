package bruteforce

import (
	"geoindexing_comparison/backend/generator"
	"testing"

	"github.com/stretchr/testify/assert"
)

var points = generator.DefaultGenerator.Points(&generator.DefaultInput, 1000)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	t.Parallel()

	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	items, _ := collection.KNNTimed(origin, 900)
	assert.Len(t, items, 900)
}

func TestUnit_PointGenerator_RadiusSearch_Ok(t *testing.T) {
	t.Parallel()

	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	collection.RangeSearchTimed(
		origin,
		(generator.DefaultInput.LatUpperBound-generator.DefaultInput.LatLowerBound)/6,
	)
}
