package generator

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DefaultNormalGenerator = ClusterGenerator{ClusterN: 6, rng: helpers.RNG()} //nolint: gochecknoglobals // Allowed here

func TestUnit_NormalGenerator_Cluster_Ok(t *testing.T) {
	t.Parallel()

	points := DefaultNormalGenerator.cluster(geo.NewPoint(55.756739, 37.627652), 10_000)
	assert.Equal(t, len(points), 10_000)
}

func TestUnit_NormalGenerator_Points_Ok(t *testing.T) {
	t.Parallel()

	points := DefaultNormalGenerator.Points(&DefaultInput, 25_000)
	assert.Equal(t, len(points), 25_000)
}
