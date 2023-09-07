package rtree

import (
	"geoindexing_comparison/geo"
	"github.com/dhconnelly/rtreego"
)

type CollectionRTree struct {
	impl   rtreego.Rtree
	points geo.Points
}

func New() *CollectionRTree {
	return &CollectionRTree{impl: *rtreego.NewTree(2, 25, 50)}
}

func (r *CollectionRTree) Name() string {
	return "RTree"
}

func (r *CollectionRTree) FromArray(points geo.Points) {
	for _, point := range points {
		r.impl.Insert(point)
	}

	r.points = points
}

func (r *CollectionRTree) Points() geo.Points {
	return r.points
}

func (r *CollectionRTree) Remove(point geo.Point) {
	r.impl.Delete(point)

	go r.points.Delete(point.ID)
}

func (r *CollectionRTree) KNN(point geo.Point, n int) geo.Points {
	spatials := r.impl.NearestNeighbors(n, []float64{point.Lat, point.Lon})
	result := make(geo.Points, len(spatials))
	for idx, spatial := range spatials {
		result[idx] = spatial.(geo.Point)
	}
	return result
}

func (r *CollectionRTree) RangeSearch(point geo.Point, radius float64) geo.Points {
	//r.impl.
	//TODO implement me
	panic("implement me")
}

func (r *CollectionRTree) String() string {
	//TODO implement me
	return r.impl.String()
}

func (r *CollectionRTree) Insert(point geo.Point) {
	r.impl.Insert(point)
}
