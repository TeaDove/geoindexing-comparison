package generator

import (
	"geoindexing_comparison/backend/geo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_NormalGenerator_Cluster_Ok(t *testing.T) {
	t.Parallel()

	DefaultNormalGenerator.cluster(geo.NewPoint(55.756739, 37.627652), 10_000)
}

func TestUnit_NormalGenerator_Points_Ok(t *testing.T) {
	t.Parallel()
	//t.Skip("Fails because normal generator works bad")

	points := DefaultNormalGenerator.Points(&DefaultInput, 25_000)
	assert.Len(t, points, 25_000)
}
