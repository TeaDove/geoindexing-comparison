package bruteforce

import (
	"time"

	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/geo"
	"golang.org/x/exp/slices"
)

type CollectionBruteforce struct {
	impl geo.Points
}

func New() addapter.Collection {
	r := CollectionBruteforce{}
	return &r
}

func (r *CollectionBruteforce) Name() string {
	return "Bruteforce"
}

func (r *CollectionBruteforce) FromArray(points geo.Points) {
	r.impl = points
}

func (r *CollectionBruteforce) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()
	r.impl = append(r.impl, point)

	return time.Now().Sub(t0)
}

func (r *CollectionBruteforce) RangeSearchTimed(
	point geo.Point,
	radius float64,
) (geo.Points, time.Duration) {
	t0 := time.Now()
	points := make(geo.Points, 0, 10)

	for _, indexPoint := range r.impl {
		if indexPoint.Lat < point.Lat+radius &&
			indexPoint.Lat > point.Lat-radius &&
			indexPoint.Lon < point.Lon+radius &&
			indexPoint.Lon > point.Lon-radius {
			points = append(points, indexPoint)
		}
	}
	dur := time.Now().Sub(t0)

	return points, dur
}

func (r *CollectionBruteforce) KNNTimed(point geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()
	if n > len(r.impl) {
		return r.impl, time.Since(t0)
	}

	type dist struct {
		idx  int
		dist float64
	}

	knnMatrix := make([]dist, 0, len(r.impl))
	for idx, indexPoint := range r.impl {
		knnMatrix = append(knnMatrix, dist{idx: idx, dist: indexPoint.DistanceTo(point)})
	}

	slices.SortFunc(knnMatrix, func(a, b dist) int {
		if a.dist < b.dist {
			return -1
		}

		if a.dist > b.dist {
			return 1
		}

		return 0
	})

	result := make(geo.Points, n)

	for idx := range n {
		result[idx] = r.impl[knnMatrix[idx].idx]
	}

	dur := time.Since(t0)

	return result, dur
}

func (r *CollectionBruteforce) String() string {
	return r.impl.String()
}
