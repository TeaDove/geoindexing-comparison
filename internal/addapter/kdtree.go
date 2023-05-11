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

func (r *CollectionKDTree) FromArray(points []geo.Point) {
	var kdPoints []kdtree.Point
	for _, point_ := range points {
		kdPoints = append(kdPoints, point_)
	}

	r.impl = kdtree.New(kdPoints)
}

func (r *CollectionKDTree) Insert(point geo.Point) {
	r.impl.Insert(point)
}

func (r *CollectionKDTree) Remove(point geo.Point) {
	r.impl.Remove(point)
}

func (r *CollectionKDTree) RangeSearch(ran float64) []geo.Point {
	_ = r.impl.RangeSearch(kdrange.New(ran))
	return []geo.Point{}
}

func (r *CollectionKDTree) String() string {
	return r.impl.String()
}
