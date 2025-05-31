package generator

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/helpers"
	"math/rand/v2"
)

type ClusterGenerator struct {
	ClusterN int

	rng *rand.Rand
}

func (r *ClusterGenerator) Points(input *Input, amount int) geo.Points {
	amountPerCluster := make([]int, r.ClusterN)
	maxPointsPerCluster := int(float64(amount/r.ClusterN) * 2)
	remain := amount

	var idx int
	for {
		currentAmount := helpers.RNG().IntN(maxPointsPerCluster)
		if remain <= currentAmount {
			amountPerCluster[idx] = remain
			break
		}

		amountPerCluster[idx] = currentAmount

		remain -= currentAmount

		idx++
		idx = idx % r.ClusterN
	}

	points := make(geo.Points, 0)
	for idx = range r.ClusterN {
		newCluster := r.cluster(
			geo.NewPoint(
				randFloat(r.rng, input.LatLowerBound, input.LatUpperBound),
				randFloat(r.rng, input.LonLowerBound, input.LonUpperBound),
			),
			amountPerCluster[idx],
		)
		points = append(points, newCluster...)
	}

	return points
}

func (r *ClusterGenerator) cluster(center geo.Point, amount int) geo.Points {
	points := make(geo.Points, amount)

	for i := range amount {
		points[i] = geo.NewPoint(
			r.rng.NormFloat64()*0.05/2+center.Lat,
			r.rng.NormFloat64()*0.05+center.Lon)
	}

	return points
}

func (r *ClusterGenerator) Point(input *Input) geo.Point {
	return RandomPoint(r.rng, input)
}
