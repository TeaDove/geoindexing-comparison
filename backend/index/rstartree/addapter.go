package rstartree

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"time"

	rtreego "github.com/anayks/go-rstar-tree"
)

type CollectionRTree struct {
	impl rtreego.Rtree
}

func New() index.IndexImpl {
	return &CollectionRTree{impl: *rtreego.NewTree(2, 1000, 100_000)}
}

func (r *CollectionRTree) FromArray(points geo.Points) {
	r.impl = *rtreego.NewTree(2, len(points), len(points)+len(points)/4)
	for _, point := range points {
		r.impl.Insert(point.ToPointForRStarTree())
	}
}

func (r *CollectionRTree) KNNTimed(point geo.Point, n uint64) (geo.Points, time.Duration) {
	t0 := time.Now()
	spatials := r.impl.NearestNeighbors(int(n), []float64{point.Lat, point.Lon})
	dur := time.Since(t0)

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
	rect, _ := rtreego.NewRect([]float64{point.Lat, point.Lon}, []float64{2 * radius, 2 * radius})
	t0 := time.Now()
	points := r.impl.SearchIntersect(rect)
	dur := time.Since(t0)

	geoPoints := make(geo.Points, len(points))
	for idx, foundPoint := range points {
		geoPoints[idx] = foundPoint.(geo.PointForRStarTree).ToPoint()
	}

	return geoPoints, dur
}

func (r *CollectionRTree) String() string {
	// TODO implement me
	return r.impl.String()
}

func (r *CollectionRTree) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.impl.Insert(point.ToPointForRStarTree())

	return time.Since(t0)
}
