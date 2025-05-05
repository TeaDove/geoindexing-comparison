package rtree

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"time"

	"github.com/dhconnelly/rtreego"
)

type CollectionRTree struct {
	impl rtreego.Rtree
}

func New() index.Impl {
	return &CollectionRTree{impl: *rtreego.NewTree(2, 1000, 100_000)}
}

func (r *CollectionRTree) FromArray(points geo.Points) {
	r.impl = *rtreego.NewTree(2, len(points), len(points)+len(points)/4)
	for _, point := range points {
		r.impl.Insert(point)
	}
}

func (r *CollectionRTree) ToArray() geo.Points {
	var res geo.Points
	for _, point := range r.impl.NearestNeighbors(r.impl.Size(), rtreego.Point([]float64{0, 0})) {
		res = append(res, point.(geo.Point))
	}

	return res
}

func (r *CollectionRTree) KNNTimed(origin geo.Point, n uint64) (geo.Points, time.Duration) {
	t0 := time.Now()
	spatials := r.impl.NearestNeighbors(int(n), []float64{origin.Lat, origin.Lon})
	dur := time.Since(t0)

	result := make(geo.Points, len(spatials))
	for idx, spatial := range spatials {
		result[idx] = spatial.(geo.Point)
	}

	return result, dur
}

func (r *CollectionRTree) RangeSearchTimed(
	origin geo.Point,
	radius float64,
) (geo.Points, time.Duration) {
	rect, _ := rtreego.NewRect([]float64{origin.Lat, origin.Lon}, []float64{2 * radius, 2 * radius})
	t0 := time.Now()
	points := r.impl.SearchIntersect(rect)
	dur := time.Since(t0)

	geoPoints := make(geo.Points, len(points))
	for idx, foundPoint := range points {
		geoPoints[idx] = foundPoint.(geo.Point)
	}

	return geoPoints, dur
}

func (r *CollectionRTree) String() string {
	// TODO implement me
	return r.impl.String()
}

func (r *CollectionRTree) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.impl.Insert(point)

	return time.Since(t0)
}
