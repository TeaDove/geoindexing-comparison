package rstartree

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/index"
	"time"

	"github.com/pkg/errors"

	rtreego "github.com/anayks/go-rstar-tree"
)

type Index struct {
	impl rtreego.Rtree
}

func New() index.Impl {
	return &Index{impl: *rtreego.NewTree(2, 1000, 100_000)}
}

func (r *Index) FromArray(points geo.Points) {
	r.impl = *rtreego.NewTree(2, len(points), len(points)+len(points)/4)
	for _, point := range points {
		r.impl.Insert(point.ToPointForRStarTree())
	}
}

func (r *Index) ToArray() geo.Points {
	var res geo.Points
	for _, point := range r.impl.NearestNeighbors(r.impl.Size(), []float64{0, 0}) {
		res = append(res, point.(geo.PointForRStarTree).ToPoint())
	}

	return res
}

func (r *Index) KNNTimed(origin geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()
	spatials := r.impl.NearestNeighbors(n, []float64{origin.Lat, origin.Lon})
	dur := time.Since(t0)

	result := make(geo.Points, len(spatials))
	for idx, spatial := range spatials {
		result[idx] = spatial.(geo.PointForRStarTree).ToPoint()
	}

	return result, dur
}

func (r *Index) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	rect, err := rtreego.NewRectFromPoints(
		[]float64{bottomLeft.Lat, bottomLeft.Lon},
		[]float64{upperRight.Lat, upperRight.Lon},
	)
	if err != nil {
		panic(errors.Wrap(err, "could not create rectangle"))
	}

	t0 := time.Now()
	points := r.impl.SearchIntersect(rect)
	dur := time.Since(t0)

	geoPoints := make(geo.Points, len(points))
	for idx, foundPoint := range points {
		geoPoints[idx] = foundPoint.(geo.PointForRStarTree).ToPoint()
	}

	return geoPoints, dur
}

func (r *Index) String() string {
	// TODO implement me
	return r.impl.String()
}

func (r *Index) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.impl.Insert(point.ToPointForRStarTree())

	return time.Since(t0)
}
