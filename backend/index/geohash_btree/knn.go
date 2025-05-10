package geohash_btree

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/geo/geohash_utils"
	"time"
)

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

	return points.GetClosestViaSort(origin, int(n)), time.Since(t0)
}
