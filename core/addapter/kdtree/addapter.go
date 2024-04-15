package kdtree

import (
	"time"

	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/geo"

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

	return time.Now().Sub(t0)
}

func (r *CollectionKDTree) RangeSearchTimed(
	point geo.Point,
	radius float64,
) (geo.Points, time.Duration) {
	t0 := time.Now()
	res := r.impl.RangeSearch(
		kdrange.New(point.Lat-radius, point.Lat+radius, point.Lon-radius, point.Lon+radius),
	)
	dur := time.Now().Sub(t0)

	return toConcrete(res), dur
}

func (r *CollectionKDTree) KNNTimed(point geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()
	res := r.impl.KNN(&point, n)
	dur := time.Now().Sub(t0)

	return toConcrete(res), dur
}

func (r *CollectionKDTree) String() string {
	return r.impl.String()
}
