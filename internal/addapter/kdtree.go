package addapter

import (
	"geoindexing_comparison/geo"
	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/kdrange"
)

type CollectionKDTree struct {
	impl *kdtree.KDTree
}

func (r *CollectionKDTree) Empty() {
	r.impl = kdtree.New([]kdtree.Point{})
}

func (r *CollectionKDTree) FromArray(points geo.Points) {
	var kdPoints []kdtree.Point
	for _, point_ := range points {
		kdPoints = append(kdPoints, point_)
	}

	r.impl = kdtree.New(kdPoints)
}

func toConcrete(pointsInterface []kdtree.Point) geo.Points {
	result := make([]geo.Point, len(pointsInterface))
	for idx, pointInterface := range pointsInterface {
		result[idx] = pointInterface.(geo.Point)
	}
	return result
}

func (r *CollectionKDTree) Points() geo.Points {
	return toConcrete(r.impl.Points())
}

func (r *CollectionKDTree) Insert(point geo.Point) {
	r.impl.Insert(point)
}

func (r *CollectionKDTree) Remove(point geo.Point) {
	r.impl.Remove(point)
}

func (r *CollectionKDTree) RangeSearch(radius float64) geo.Points {
	return toConcrete(r.impl.RangeSearch(kdrange.New(radius)))
}

func (r *CollectionKDTree) KNN(point geo.Point, n int) geo.Points {
	return toConcrete(r.impl.KNN(point, n))
}

func (r *CollectionKDTree) String() string {
	return r.impl.String()
}
