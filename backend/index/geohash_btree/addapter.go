package geohash_btree

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"github.com/tidwall/btree"
	"time"
)

type CollectionGeohash struct {
	btree       btree.Map[uint64, geo.Points]
	geohashBits uint
}

func Factory(geohashPrecisionChars uint) func() index.Impl {
	return func() index.Impl {
		return &CollectionGeohash{btree: *btree.NewMap[uint64, geo.Points](2), geohashBits: geohashPrecisionChars * 5}
	}
}

func (r *CollectionGeohash) geohash(point geo.Point) uint64 {
	return point.Geohash(r.geohashBits)
}

func (r *CollectionGeohash) FromArray(points geo.Points) {
	for _, point := range points {
		r.Insert(point)
	}
}

func (r *CollectionGeohash) ToArray() geo.Points {
	var points geo.Points
	for _, arr := range r.btree.Values() {
		points = append(points, arr...)
	}
	return points
}

func (r *CollectionGeohash) Insert(point geo.Point) {
	v := r.geohash(point)
	points, _ := r.btree.Get(v)
	points = append(points, point)

	r.btree.Set(r.geohash(point), points)
}

func (r *CollectionGeohash) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()
	r.Insert(point)
	return time.Since(t0)
}

func (r *CollectionGeohash) String() string {
	points := r.ToArray()
	return points.String()
}
