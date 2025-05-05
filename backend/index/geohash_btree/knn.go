package geohash_btree

import (
	"geoindexing_comparison/backend/geo"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/mmcloughlin/geohash"
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

func (r *CollectionGeohash) findInNeighbors(processedNeighbors mapset.Set[uint64], foundPoints *geo.Points, n int) {
	for neighbor := range processedNeighbors.Each {
		neighbors := geohash.NeighborsIntWithPrecision(neighbor, r.geohashPrecision)
		for _, foundNeighbor := range neighbors {
			if processedNeighbors.Contains(foundNeighbor) {
				continue
			}

			points, _ := r.btree.Get(neighbor)
			*foundPoints = append(*foundPoints, points...)
			if len(points) >= n {
				return
			}

			processedNeighbors.Add(neighbor)
		}
	}
}

func (r *CollectionGeohash) KNNTimed(origin geo.Point, n uint64) (geo.Points, time.Duration) {
	t0 := time.Now()

	originGeohash := r.geohash(origin)
	points, _ := r.btree.Get(originGeohash)
	processedGeohashes := mapset.NewSet[uint64]()
	processedGeohashes.Add(originGeohash)

	for len(points) < int(n) {
		r.findInNeighbors(processedGeohashes, &points, int(n))
	}

	_, maxDistance := findMostDistantPoint(origin, points)
	pointsInRange := r.rangeSearch(origin, maxDistance)
	pointsInRange.SortByDistance(origin)

	return pointsInRange[:n], time.Since(t0)
}
