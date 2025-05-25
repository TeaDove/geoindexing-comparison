package generator

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/helpers"
	"math/rand/v2"
)

type NormalGenerator struct {
	ClusterN int

	rng *rand.Rand
}

var DefaultNormalGenerator = NormalGenerator{ClusterN: 6, rng: helpers.RNG()} //nolint: gochecknoglobals // Allowed here

func (r *NormalGenerator) Points(input *Input, amount int) geo.Points {
	amounts := make([]int64, r.ClusterN)
	mapPerCluster := int64(float64(int(amount)/r.ClusterN) * 0.2)
	remain := int64(amount)

	for idx := range amounts {
		amounts[idx] = helpers.RNG().Int64N(mapPerCluster)

		remain -= amounts[idx]
		if remain <= 0 {
			break
		}
	}

	points := make(geo.Points, 0)
	for idx := range r.ClusterN {
		points = append(points, r.cluster(
			geo.NewPoint(
				randFloat(r.rng, input.LatLowerBound, input.LatUpperBound),
				randFloat(r.rng, input.LonLowerBound, input.LonUpperBound),
			),
			int(amounts[idx]),
		)...,
		)
	}

	return points
}

func (r *NormalGenerator) cluster(center geo.Point, amount int) geo.Points {
	points := make(geo.Points, amount)

	for i := range amount {
		points[i] = geo.NewPoint(
			r.rng.NormFloat64()*0.05/2+center.Lat,
			r.rng.NormFloat64()*0.05+center.Lon)
	}

	return points
}

func (r *NormalGenerator) Point(input *Input) geo.Point {
	return geo.NewPoint(
		randFloat(r.rng, input.LatLowerBound, input.LatUpperBound),
		randFloat(r.rng, input.LonLowerBound, input.LonUpperBound),
	)
}
