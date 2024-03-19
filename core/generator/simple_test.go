package generator

import (
	"testing"

	"geoindexing_comparison/core/geo"

	"github.com/guregu/null"

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

	assert.Equal(t, 10_000, len(points))
}

func TestUnit_PointGenerator_MustExport_Ok(t *testing.T) {
	t.Parallel()

	points := DefaultGenerator.Points(&DefaultInput, 10_000)
	points.
		ToPointExtended().
		MustExport(&geo.ExportInput{Filename: null.StringFrom("../../data/test-points.csv")})
}
