package geobtree

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/geo/geohash_utils"
	"time"
)

func (r *Index) KNNTimed(origin geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()

	originGeohash := r.geohash(origin)

	var points geo.Points

	for neighbors := range geohash_utils.NeighborIterSquared(originGeohash, r.geohashBits) {
		points = append(points, r.getMany(neighbors)...)
		if len(points) >= n {
			break
		}
	}

	mostDistance := points.FindMostDistant(origin, r.metric)
	bottomLeft := origin.AddLatitude(-mostDistance).AddLongitude(-mostDistance)
	upperRight := origin.AddLatitude(mostDistance).AddLongitude(mostDistance)
	bboxedPoints := r.bbox(bottomLeft, upperRight)

	return bboxedPoints.GetClosestViaSort(origin, n), time.Since(t0)
}
