package geohash_btree

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/geo/geohash_utils"
	"time"
)

func (r *CollectionGeohash) getMany(hashed []uint64) geo.Points {
	var (
		points      geo.Points
		foundPoints geo.Points
	)

	for _, hash := range hashed {
		foundPoints, _ = r.btree.Get(hash)
		points = append(points, foundPoints...)
	}

	return points
}

func (r *CollectionGeohash) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	t0 := time.Now()

	return r.bbox(bottomLeft, upperRight), time.Since(t0)
}

func (r *CollectionGeohash) bbox(bottomLeft geo.Point, upperRight geo.Point) geo.Points {
	bbox := geohash_utils.NewBBox(bottomLeft.Lat, bottomLeft.Lon, upperRight.Lat, upperRight.Lon, r.geohashBits)
	points := r.getMany(bbox.Inner())
	outerPoints := r.getMany(bbox.Perimeter())

	for _, point := range outerPoints {
		if point.InsideBBox(bottomLeft, upperRight) {
			points = append(points, point)
		}
	}

	return points
}
