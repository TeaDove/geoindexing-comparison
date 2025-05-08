package geohash_btree

import (
	"geoindexing_comparison/backend/geo"
	"time"
)

func findMostDistantPoint(origin geo.Point, points geo.Points) (geo.Point, float64) {
	var (
		maxDistance      float64
		distance         float64
		mostDistantPoint geo.Point
	)
	for _, point := range points {
		distance = point.DistanceTo(origin)
		if distance > maxDistance {
			maxDistance = point.DistanceTo(origin)
			mostDistantPoint = point
		}
	}

	return mostDistantPoint, maxDistance
}

func (r *CollectionGeohash) KNNTimed(origin geo.Point, n uint64) (geo.Points, time.Duration) {
	t0 := time.Now()

	originGeohash := r.geohash(origin)
	iters := 0
	points, _ := r.btree.Get(originGeohash)

	for neighbor := range geo.GeohashNeighborIter(originGeohash, r.geohashBits) {
		neighborPoints, _ := r.btree.Get(neighbor)
		points = append(points, neighborPoints...)

		iters++
		if len(neighborPoints) != 0 {
			println(neighbor, len(points), n, len(points) >= int(n), len(neighborPoints), iters)
		}

		if len(points) >= int(n) {
			break
		}
	}

	_, maxDistance := findMostDistantPoint(origin, points)
	pointsInRange := r.rangeSearch(origin, maxDistance)

	return pointsInRange.GetClosestViaSort(origin, int(n)), time.Since(t0)
}
