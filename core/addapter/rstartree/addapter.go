package rstartree

import (
	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/geo"
	"github.com/anayks/go-rstar-tree"
	"time"
)

type CollectionRTree struct {
	impl rtreego.Rtree
}

func New() addapter.Collection {
	return &CollectionRTree{impl: *rtreego.NewTree(2, 25, 50)}
}

func (r *CollectionRTree) Name() string {
	return "RStarTree"
}

func (r *CollectionRTree) FromArray(points geo.Points) {
	for _, point := range points {
		r.impl.Insert(point.ToPointForRStarTree())
	}
}

func (r *CollectionRTree) KNNTimed(point geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()
	spatials := r.impl.NearestNeighbors(n, []float64{point.Lat, point.Lon})
	dur := time.Now().Sub(t0)

	result := make(geo.Points, len(spatials))
	for idx, spatial := range spatials {
		result[idx] = spatial.(geo.PointForRStarTree).ToPoint()
	}
	return result, dur
}

func (r *CollectionRTree) RangeSearchTimed(
	point geo.Point,
	radius float64,
) (geo.Points, time.Duration) {
	t0 := time.Now()
	rect, _ := rtreego.NewRect([]float64{point.Lat, point.Lon}, []float64{radius, radius})
	points := r.impl.SearchIntersect(rect)
	dur := time.Now().Sub(t0)

	geoPoints := make(geo.Points, len(points))
	for idx, foundPoint := range points {
		geoPoints[idx] = foundPoint.(geo.PointForRStarTree).ToPoint()
	}
	return geoPoints, dur
}

func (r *CollectionRTree) String() string {
	//TODO implement me
	return r.impl.String()
}

func (r *CollectionRTree) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()
	r.impl.Insert(point.ToPointForRStarTree())
	return time.Now().Sub(t0)
}