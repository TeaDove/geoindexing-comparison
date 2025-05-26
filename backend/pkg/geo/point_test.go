package geo

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMeters(t *testing.T) {
	t.Parallel()

	origin := NewPoint(55.803459, 37.798224)
	moved := origin.AddLatitude(0.1).AddLongitude(0.1)

	assert.InEpsilon(t, math.Sqrt(200)/100, moved.DistanceHaversine(origin), 0.00001)

	moved = origin.AddLatitude(0.3)
	assert.InEpsilon(t, 0.3, moved.DistanceHaversine(origin), 0.00001)

	moved = origin.AddLongitude(0.3)
	assert.InEpsilon(t, 0.3, moved.DistanceHaversine(origin), 0.00001)
}
