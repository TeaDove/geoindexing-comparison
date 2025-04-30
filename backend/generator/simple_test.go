package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	t.Parallel()

	testPoint := DefaultGenerator.Point(&DefaultInput)

	assert.LessOrEqual(t, DefaultInput.LatLowerBound, testPoint.Lat)
	assert.GreaterOrEqual(t, DefaultInput.LatUpperBound, testPoint.Lat)
	assert.LessOrEqual(t, DefaultInput.LonLowerBound, testPoint.Lon)
	assert.GreaterOrEqual(t, DefaultInput.LonUpperBound, testPoint.Lon)
}

func TestUnit_PointGenerator_GeneratePoints_Ok(t *testing.T) {
	t.Parallel()

	points := DefaultGenerator.Points(&DefaultInput, 10_000)

	assert.Len(t, points, 10_000)
}

func TestUnit_PointGenerator_MustExport_Ok(t *testing.T) {
	t.Parallel()

	DefaultGenerator.Points(&DefaultInput, 10_000)
}
