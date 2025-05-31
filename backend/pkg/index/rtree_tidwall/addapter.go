package rtree_tidwall

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/index"
	"github.com/tidwall/rtree"
	"time"
)

type Index struct {
	impl rtree.RTreeG[geo.Point]
}

func New() index.Impl {
	return &Index{impl: rtree.RTreeG[geo.Point]{}}
}

func (r *Index) FromArray(points geo.Points) {
	for _, point := range points {
		r.insert(point)
	}
}

func (r *Index) ToArray() geo.Points {
	var points geo.Points
	r.impl.Scan(func(_, _ [2]float64, data geo.Point) bool {
		points = append(points, data)
		return true
	})

	return points
}

func (r *Index) insert(point geo.Point) {
	r.impl.Insert([2]float64{point.Lon, point.Lat}, [2]float64{point.Lon, point.Lat}, point)
}

func (r *Index) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()
	r.insert(point)
	return time.Since(t0)
}

func (r *Index) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	t0 := time.Now()
	var points geo.Points
	r.impl.Search([2]float64{bottomLeft.Lon, bottomLeft.Lat}, [2]float64{upperRight.Lon, upperRight.Lat},
		func(_, _ [2]float64, data geo.Point) bool {
			points = append(points, data)
			return true
		})

	return points, time.Since(t0)
}

func (r *Index) KNNTimed(origin geo.Point, k int) (geo.Points, time.Duration) {
	t0 := time.Now()

	var points geo.Points
	r.impl.Nearby(
		rtree.BoxDist[float64, geo.Point]([2]float64{origin.Lon, origin.Lat}, [2]float64{origin.Lon, origin.Lat}, nil),
		func(_, _ [2]float64, data geo.Point, dist float64) bool {
			points = append(points, data)
			if len(points) >= k {
				return false
			}
			return true
		},
	)

	return points, time.Since(t0)
}
