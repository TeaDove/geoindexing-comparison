package kdtree

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/index"
	"time"

	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/kdrange"
)

type Index struct {
	impl kdtree.KDTree
}

func New() index.Impl {
	r := Index{}
	r.impl = *kdtree.New([]kdtree.Point{})

	return &r
}

func (r *Index) FromArray(points geo.Points) {
	kdPoints := make([]kdtree.Point, len(points))
	for idx, point := range points {
		kdPoints[idx] = &point
	}

	r.impl = *kdtree.New(kdPoints)
}

func (r *Index) ToArray() geo.Points {
	var res geo.Points
	for _, point := range r.impl.Points() {
		res = append(res, *point.(*geo.Point))
	}

	return res
}

func toConcrete(pointsInterface []kdtree.Point) geo.Points {
	result := make([]geo.Point, len(pointsInterface))

	for idx, pointInterface := range pointsInterface {
		switch r := pointInterface.(type) {
		// TODO remove this poor taste solution
		case geo.Point:
			result[idx] = r
		case *geo.Point:
			result[idx] = *r
		default:
			panic("invalid type")
		}
	}

	return result
}

func (r *Index) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.impl.Insert(&point)

	return time.Since(t0)
}

func (r *Index) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	t0 := time.Now()
	res := r.impl.RangeSearch(
		kdrange.New(bottomLeft.Lat, upperRight.Lat, bottomLeft.Lon, upperRight.Lon),
	)
	dur := time.Since(t0)

	return toConcrete(res), dur
}

func (r *Index) KNNTimed(origin geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()
	res := r.impl.KNN(&origin, n)
	dur := time.Since(t0)

	return toConcrete(res), dur
}

func (r *Index) String() string {
	return r.impl.String()
}
