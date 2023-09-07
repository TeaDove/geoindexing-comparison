package cases

import (
	"geoindexing_comparison/addapter"
	"geoindexing_comparison/addapter/kdtree"
	"geoindexing_comparison/addapter/rtree"
	"geoindexing_comparison/geo"
)

type Case struct {
	Collection addapter.Collection
}

func NewCollections() []addapter.Collection {
	return []addapter.Collection{
		rtree.New(),
		kdtree.New(),
	}
}

type KNNInput struct {
	Amount int
	Points geo.Points
	Origin geo.Point
}
