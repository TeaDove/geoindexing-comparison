package geobrtree

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/geo/geohash_utils"
	"github.com/tidwall/rtree"
	"time"
)

func (r *Index) getMany(hashed []uint64) []*rtree.RTreeG[geo.Point] {
	var (
		points    []*rtree.RTreeG[geo.Point]
		foundTree *rtree.RTreeG[geo.Point]
		ok        bool
	)

	for _, hash := range hashed {
		foundTree, ok = r.btree.Get(hash)
		if ok {
			points = append(points, foundTree)
		}

	}

	return points
}

func (r *Index) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	t0 := time.Now()

	return r.bbox(bottomLeft, upperRight), time.Since(t0)
}

func (r *Index) bbox(bottomLeft geo.Point, upperRight geo.Point) geo.Points {
	bbox := geohash_utils.NewBBox(bottomLeft.Lat, bottomLeft.Lon, upperRight.Lat, upperRight.Lon, r.geohashBits)

	var points geo.Points
	innerTrees := r.getMany(bbox.Inner())
	for _, tree := range innerTrees {
		tree.Scan(func(_, _ [2]float64, data geo.Point) bool {
			points = append(points, data)
			return true
		})
	}

	outerTrees := r.getMany(bbox.Perimeter())

	for _, tree := range outerTrees {
		tree.Search(
			[2]float64{bottomLeft.Lon, bottomLeft.Lat},
			[2]float64{upperRight.Lon, upperRight.Lat},
			func(min, max [2]float64, data geo.Point) bool {
				points = append(points, data)
				return true
			})
	}

	return points
}
