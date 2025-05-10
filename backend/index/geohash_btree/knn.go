package geohash_btree

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/geo/geohash_utils"
	"time"
)

func findMostDistant(origin geo.Point, points geo.Points) float64 {
	var (
		mostDistance float64
		distance     float64
	)

	for _, point := range points {
		distance = point.DistanceTo(origin)
		if distance > mostDistance {
			mostDistance = distance
		}
	}

	return mostDistance
}

func (r *CollectionGeohash) KNNTimed(origin geo.Point, n uint64) (geo.Points, time.Duration) {
	t0 := time.Now()

	originGeohash := r.geohash(origin)
	var points geo.Points
	for neighbors := range geohash_utils.NeighborIterSquared(originGeohash, r.geohashBits) {
		points = append(points, r.getMany(neighbors)...)
		if len(points) >= int(n) {
			break
		}
	}

	mostDistance := findMostDistant(origin, points)
	bottomLeft := origin.AddLatitude(-mostDistance).AddLongitude(-mostDistance)
	upperRight := origin.AddLatitude(mostDistance).AddLongitude(mostDistance)
	bboxedPoints := r.bbox(bottomLeft, upperRight)

	return bboxedPoints.GetClosestViaSort(origin, int(n)), time.Since(t0)
}
