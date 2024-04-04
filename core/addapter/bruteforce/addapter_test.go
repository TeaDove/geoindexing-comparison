package bruteforce

import (
	"github.com/stretchr/testify/assert"
	"testing"

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
