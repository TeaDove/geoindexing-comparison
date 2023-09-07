package kdtree

import (
	"geoindexing_comparison/geo"

	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/kdrange"
)

type CollectionKDTree struct {
	impl kdtree.KDTree
}

func New() *CollectionKDTree {
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

func (r *CollectionKDTree) Points() geo.Points {
	return toConcrete(r.impl.Points())
}

func (r *CollectionKDTree) Insert(point geo.Point) {
	r.impl.Insert(&point)
}

func (r *CollectionKDTree) Remove(point geo.Point) {
	r.impl.Remove(&point)
}

func (r *CollectionKDTree) RangeSearch(point geo.Point, radius float64) geo.Points {
	return toConcrete(
		r.impl.RangeSearch(
			kdrange.New(point.Lat-radius, point.Lat+radius, point.Lon-radius, point.Lon+radius),
		),
	)
}

func (r *CollectionKDTree) KNN(point geo.Point, n int) geo.Points {
	return toConcrete(r.impl.KNN(&point, n))
}

func (r *CollectionKDTree) String() string {
	return r.impl.String()
}
