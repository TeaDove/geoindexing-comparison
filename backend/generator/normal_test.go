package generator

import (
	"geoindexing_comparison/backend/geo"
	"testing"
)

func TestUnit_NormalGenerator_Cluster_Ok(t *testing.T) {
	t.Parallel()

	DefaultNormalGenerator.cluster(geo.NewPoint(55.756739, 37.627652), 10_000)
}

func TestUnit_NormalGenerator_Points_Ok(t *testing.T) {
	t.Parallel()

	DefaultNormalGenerator.Points(&DefaultInput, 25_000)
}
