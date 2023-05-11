package generator

import (
	"geoindexing_comparison/geo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	testPoint := DefaultGenerator.GeneratePoint()

	assert.LessOrEqual(t, DefaultGenerator.LatLowerBound, testPoint.Lat)
	assert.GreaterOrEqual(t, DefaultGenerator.LatUpperBound, testPoint.Lat)
	assert.LessOrEqual(t, DefaultGenerator.LonLowerBound, testPoint.Lon)
	assert.GreaterOrEqual(t, DefaultGenerator.LonUpperBound, testPoint.Lon)
}

func TestUnit_PointGenerator_GeneratePoints_Ok(t *testing.T) {
	points := DefaultGenerator.GeneratePointsDefaultAmount()

	assert.Equal(t, DefaultGenerator.PointsAmount, len(points))
}

func TestUnit_PointGenerator_MustExport_Ok(t *testing.T) {
	points := DefaultGenerator.GeneratePoints(DefaultGenerator.PointsAmount)
	geo.MustExport(points)
}
