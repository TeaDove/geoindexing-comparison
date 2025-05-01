package generator

import (
	"geoindexing_comparison/backend/geo"
	"math/rand/v2"
)

type NormalGenerator struct {
	ClusterN int
}

var DefaultNormalGenerator = NormalGenerator{ClusterN: 6} //nolint: gochecknoglobals // Allowed here

func (r *NormalGenerator) Points(input *Input, amount uint64) geo.Points {
	rng := rand.New(rand.NewPCG(0, 0)) //nolint: gosec // Allowed here

	amounts := make([]int64, r.ClusterN)
	mapPerCluster := int64(float64(int(amount)/r.ClusterN) * 0.2)
	remain := int64(amount)

	for idx := range amounts {
		amounts[idx] = rng.Int64N(mapPerCluster)

		remain -= amounts[idx]
		if remain <= 0 {
			break
		}
	}

	points := make(geo.Points, 0, amount)
	for idx := range r.ClusterN {
		points = append(points, r.cluster(
			geo.NewPoint(
				randFloat(input.LatLowerBound, input.LatUpperBound),
				randFloat(input.LonLowerBound, input.LonUpperBound),
			),
			int(amounts[idx]),
		)...,
		)
	}

	return points
}

func (r *NormalGenerator) cluster(center geo.Point, amount int) geo.Points {
	rng := rand.New(rand.NewPCG(0, 0)) //nolint: gosec // Allowed here
	points := make(geo.Points, amount)

	for i := range amount {
		points[i] = geo.NewPoint(
			rng.NormFloat64()*0.05/2+center.Lat,
			rng.NormFloat64()*0.05+center.Lon)
	}

	return points
}

func (r *NormalGenerator) Point(input *Input) geo.Point {
	return geo.NewPoint(
		randFloat(input.LatLowerBound, input.LatUpperBound),
		randFloat(input.LonLowerBound, input.LonUpperBound),
	)
}
