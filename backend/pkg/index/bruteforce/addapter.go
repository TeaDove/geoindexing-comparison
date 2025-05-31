package bruteforce

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/index"
	"time"
)

type Index struct {
	impl geo.Points
}

func New() index.Impl {
	return &Index{}
}

func (r *Index) FromArray(points geo.Points) {
	r.impl = points
}

func (r *Index) ToArray() geo.Points {
	return r.impl
}

func (r *Index) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.impl = append(r.impl, point)

	return time.Since(t0)
}

func (r *Index) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	t0 := time.Now()
	points := make(geo.Points, 0, 10)

	for _, point := range r.impl {
		if point.InsideBBox(bottomLeft, upperRight) {
			points = append(points, point)
		}
	}

	dur := time.Since(t0)

	return points, dur
}

func (r *Index) KNNTimed(origin geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()
	return r.impl.GetClosestViaSort(origin, n), time.Since(t0)
}

func (r *Index) String() string {
	return r.impl.String()
}
