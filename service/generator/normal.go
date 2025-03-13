package generator

import (
	"geoindexing_comparison/service/geo"
	"time"

	rand "golang.org/x/exp/rand"
)

type NormalGenerator struct {
	ClusterN int
}

var DefaultNormalGenerator = NormalGenerator{ClusterN: 6}

func (r *NormalGenerator) Points(input *Input, amount uint64) geo.Points {
	rng := rand.New(rand.NewSource(uint64(time.Now().Unix())))

	amounts := make([]int64, r.ClusterN)
	mapPerCluster := int64(float64(int(amount)/r.ClusterN) * 0.2)
	remain := int64(amount)

	for idx := range amounts {
		amounts[idx] = rng.Int63n(mapPerCluster)

		remain -= amounts[idx]
		if remain <= 0 {
			break
		}
	}

	points := make(geo.Points, 0, amount)
	for idx := 0; idx < r.ClusterN; idx++ {
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
	rng := rand.New(rand.NewSource(uint64(time.Now().Unix())))
	points := make(geo.Points, amount)

	for i := 0; i < amount; i++ {
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
