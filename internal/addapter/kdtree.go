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

func (r *CollectionKDTree) RangeSearch(radius float64) []geo.Point {
	pointsInterface := r.impl.RangeSearch(kdrange.New(radius))
	result := make([]geo.Point, len(pointsInterface))
	for idx, pointInterface := range pointsInterface {
		result[idx] = pointInterface.(geo.Point)
	}
	return result
}

func (r *CollectionKDTree) KNN(point geo.Point, n int) []geo.Point {
	pointsInterface := r.impl.KNN(point, n)
	result := make([]geo.Point, len(pointsInterface))
	for idx, pointInterface := range pointsInterface {
		result[idx] = pointInterface.(geo.Point)
	}
	return result
}

func (r *CollectionKDTree) String() string {
	return r.impl.String()
}
