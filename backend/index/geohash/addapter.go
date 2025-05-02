package geohash

import (
	"geoindexing_comparison/backend/geo"
	"github.com/tidwall/btree"
)

type CollectionGeohash struct {
	btree            btree.Map[uint64, geo.Point]
	geohashPrecision uint
}

func New() CollectionGeohash {
	collection := CollectionGeohash{btree: *btree.NewMap[uint64, geo.Point](1), geohashPrecision: 7}

	return collection
}

func (c *CollectionGeohash) FromArray(points geo.Points) {
	for _, point := range points {
		c.Insert(point)
	}
}

func (c *CollectionGeohash) Points() geo.Points {
	return c.btree.Values()
}

func (c *CollectionGeohash) Insert(point geo.Point) {
	c.btree.Set(point.Geohash(c.geohashPrecision), point)
}

func (c *CollectionGeohash) Remove(point geo.Point) {
	c.btree.Delete(point.Geohash(c.geohashPrecision))
}

func (c *CollectionGeohash) KNN(_ geo.Point, _ int) geo.Points {
	// TODO implement me
	panic("implement me")
}

func (c *CollectionGeohash) String() string {
	// TODO implement me
	panic("implement me")
}
