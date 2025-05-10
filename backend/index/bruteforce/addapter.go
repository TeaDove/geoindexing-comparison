package bruteforce

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"time"
)

type CollectionBruteforce struct {
	impl geo.Points
}

func New() index.Impl {
	return &CollectionBruteforce{}
}

func (r *CollectionBruteforce) FromArray(points geo.Points) {
	r.impl = points
}

func (r *CollectionBruteforce) ToArray() geo.Points {
	return r.impl
}

func (r *CollectionBruteforce) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.impl = append(r.impl, point)

	return time.Since(t0)
}

func (r *CollectionBruteforce) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
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

func (r *CollectionBruteforce) KNNTimed(origin geo.Point, n uint64) (geo.Points, time.Duration) {
	t0 := time.Now()
	return r.impl.GetClosestViaSort(origin, int(n)), time.Since(t0)
}

func (r *CollectionBruteforce) String() string {
	return r.impl.String()
}
