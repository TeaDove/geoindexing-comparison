package h3_btree

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/geo/h3_utils"
	"time"
)

func (r *CollectionGeohash) KNNTimed(origin geo.Point, n uint64) (geo.Points, time.Duration) {
	t0 := time.Now()

	originHash := r.hash(origin)
	var points geo.Points
	for neighbors := range h3_utils.GridDiskInf(originHash) {
		points = append(points, r.getMany(neighbors)...)
		if len(points) >= int(n) {
			break
		}
	}

	mostDistance := points.FindMostDistant(origin, r.metric)
	bottomLeft := origin.AddLatitude(-mostDistance).AddLongitude(-mostDistance)
	upperRight := origin.AddLatitude(mostDistance).AddLongitude(mostDistance)
	bboxedPoints := r.bbox(bottomLeft, upperRight)

	return bboxedPoints.GetClosestViaSort(origin, int(n)), time.Since(t0)
}
