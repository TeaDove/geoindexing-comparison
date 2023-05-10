package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var generator = DefaultGenerator()

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	point := generator.GeneratePoint()

	assert.LessOrEqual(t, generator.LatLowerBound, point.Lat)
	assert.GreaterOrEqual(t, generator.LatUpperBound, point.Lat)
	assert.LessOrEqual(t, generator.LonLowerBound, point.Lon)
	assert.GreaterOrEqual(t, generator.LonUpperBound, point.Lon)
}

func TestUnit_PointGenerator_GeneratePoints_Ok(t *testing.T) {
	point := generator.GeneratePoints(5_000_000)

	assert.Equal(t, 5_000_000, len(point))
}
