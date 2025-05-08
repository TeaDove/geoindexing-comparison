package geohash_btree

import (
	"geoindexing_comparison/backend/geo"
	"github.com/mmcloughlin/geohash"
	"time"
)

func (r *CollectionGeohash) findNear(pointGeohash uint64, origin geo.Point, distance float64) geo.Points {
	var points geo.Points
	foundPoints, _ := r.btree.Get(pointGeohash)
	for _, point := range foundPoints {
		if origin.DistanceTo(point) <= distance {
			points = append(points, point)
		}
	}

	return points
}

func (r *CollectionGeohash) searchNeighbors(origin geo.Point, originHash uint64, radius float64, neighbors []uint64) geo.Points {

	var points geo.Points
	for _, neighbor := range neighbors {
		points = append(points, r.findNear(neighbor, origin, radius)...)
		lat, lng := geohash.DecodeIntWithPrecision(neighbor, r.geohashBits)
		if origin.DistanceToLatLng(lat, lng) <= radius {
			points = append(points, r.searchNeighbors(origin, originHash, radius, geohash.NeighborsIntWithPrecision(originHash, r.geohashBits))...)
		}
	}

	return points
}

func (r *CollectionGeohash) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	t0 := time.Now()

	return r.rangeSearch(origin, radius), time.Since(t0)
}

func (r *CollectionGeohash) rangeSearch(origin geo.Point, radius float64) geo.Points {
	originGeohash := r.geohash(origin)
	points := r.findNear(originGeohash, origin, radius)

	neighbors := geohash.NeighborsIntWithPrecision(r.geohash(origin), r.geohashBits)
	points = append(points, r.searchNeighbors(origin, originGeohash, radius, neighbors)...)

	return points
}
