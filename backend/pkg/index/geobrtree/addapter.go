package geobrtree

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/geo/distance_utils"
	"geoindexing_comparison/pkg/index"
	"github.com/tidwall/rtree"
	"time"

	"github.com/tidwall/btree"
)

type Index struct {
	btree       btree.Map[uint64, *rtree.RTreeG[geo.Point]]
	geohashBits uint
	metric      distance_utils.Metric
}

func Factory(geohashPrecisionChars uint) func() index.Impl {
	return func() index.Impl {
		return &Index{
			btree:       *btree.NewMap[uint64, *rtree.RTreeG[geo.Point]](2),
			geohashBits: geohashPrecisionChars * 5,
			metric:      distance_utils.DistanceHaversine,
		}
	}
}

func (r *Index) geohash(point geo.Point) uint64 {
	return point.Geohash(r.geohashBits)
}

func (r *Index) FromArray(points geo.Points) {
	for _, point := range points {
		r.insert(point)
	}
}

func (r *Index) ToArray() geo.Points {
	var points geo.Points
	for _, arr := range r.btree.Values() {
		arr.Scan(func(min, max [2]float64, data geo.Point) bool {
			points = append(points, data)
			return true
		})
	}

	return points
}

func (r *Index) insert(point geo.Point) {
	v := r.geohash(point)

	tree, ok := r.btree.Get(v)
	if !ok {
		tree = &rtree.RTreeG[geo.Point]{}
		r.btree.Set(v, tree)
	}

	tree.Insert([2]float64{point.Lon, point.Lat}, [2]float64{point.Lon, point.Lat}, point)
}

func (r *Index) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.insert(point)

	return time.Since(t0)
}
