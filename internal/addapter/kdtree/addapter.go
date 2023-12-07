package kdtree

import (
	"geoindexing_comparison/addapter"
	"geoindexing_comparison/geo"
	"time"

	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/kdrange"
)

type CollectionKDTree struct {
	impl kdtree.KDTree
}

func New() addapter.Collection {
	r := CollectionKDTree{}
	r.impl = *kdtree.New([]kdtree.Point{})
	return &r
}

func (r *CollectionKDTree) Name() string {
	return "KDTree"
}

func (r *CollectionKDTree) FromArray(points geo.Points) {
	kdPoints := make([]kdtree.Point, len(points))
	for idx, point_ := range points {
		kdPoints[idx] = &point_
	}

	r.impl = *kdtree.New(kdPoints)
}

func toConcrete(pointsInterface []kdtree.Point) geo.Points {
	result := make([]geo.Point, len(pointsInterface))
	for idx, pointInterface := range pointsInterface {
		switch r := pointInterface.(type) {
		// TODO remove this poor taste solution
		case geo.Point:
			result[idx] = r
		case *geo.Point:
			result[idx] = *r
		default:
			panic("invalid type")
		}
	}
	return result
}

func (r *CollectionKDTree) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()
	r.impl.Insert(&point)

	return t0.Sub(time.Now())
}

func (r *CollectionKDTree) RangeSearchTimed(point geo.Point, radius float64) (geo.Points, time.Duration) {
	t0 := time.Now()
	res := r.impl.RangeSearch(kdrange.New(point.Lat-radius, point.Lat+radius, point.Lon-radius, point.Lon+radius))
	dur := t0.Sub(time.Now())

	return toConcrete(res), dur
}

func (r *CollectionKDTree) KNNTimed(point geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()
	res := r.impl.KNN(&point, n)
	dur := t0.Sub(time.Now())

	return toConcrete(res), dur
}

func (r *CollectionKDTree) String() string {
	return r.impl.String()
}
