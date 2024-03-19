package generator

import (
	"geoindexing_comparison/core/geo"
	"github.com/guregu/null"
	"testing"
)

func TestUnit_NormalGenerator_Cluster_Ok(t *testing.T) {
	points := DefaultNormalGenerator.cluster(geo.NewPoint(55.756739, 37.627652), 10_000)

	points.
		ToPointExtended().
		MustExport(&geo.ExportInput{Filename: null.StringFrom("../../data/test-points.csv")})
}

func TestUnit_NormalGenerator_Points_Ok(t *testing.T) {
	points := DefaultNormalGenerator.Points(&DefaultInput, 25_000)

	points.
		ToPointExtended().
		MustExport(&geo.ExportInput{Filename: null.StringFrom("../../data/test-points.csv")})
}
