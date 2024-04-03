package bruteforce

import (
	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/geo"
	"time"
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

	knnMatrix := make([]float64, 0, len(r.impl))
	for _, indexPoint := range r.impl {
		knnMatrix = append(knnMatrix, indexPoint.DistanceTo(point))

	}

	result := make(geo.Points, n)
	// TODO make real knn

	//for idx, indexDistance := range knnMatrix {
	//
	//}
	//
	//res := r.impl.KNN(&point, n)
	dur := time.Since(t0)

	return result, dur
}

func (r *CollectionBruteforce) String() string {
	return r.impl.String()
}
