package geohash

import (
	"geoindexing_comparison/backend/geo"

	"github.com/dghubble/trie"
)

type CollectionGeohash struct {
	trie trie.RuneTrie
}

func New() CollectionGeohash {
	panic("Not yet implemented")

	collection := CollectionGeohash{trie: *trie.NewRuneTrie()}

	return collection
}

func (c *CollectionGeohash) FromArray(points geo.Points) {
	for _, point := range points {
		c.trie.Put(point.Geohash(), point)
	}
}

func (c *CollectionGeohash) Points() geo.Points {
	// TODO implement me
	panic("implement me")
}

func (c *CollectionGeohash) Insert(point geo.Point) {
	c.trie.Put(point.Geohash(), point)
}

func (c *CollectionGeohash) Remove(point geo.Point) {
	c.trie.Delete(point.Geohash())
	// TODO implement me
	panic("implement me")
}

func (c *CollectionGeohash) KNN(point geo.Point, n int) geo.Points {
	// TODO implement me
	panic("implement me")
}

func (c *CollectionGeohash) String() string {
	// TODO implement me
	panic("implement me")
}
