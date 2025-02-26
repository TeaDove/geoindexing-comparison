package indexes

import (
	"geoindexing_comparison/service/index"
	"geoindexing_comparison/service/index/bruteforce"
	"geoindexing_comparison/service/index/kdtree"
	"geoindexing_comparison/service/index/quadtree"
	"geoindexing_comparison/service/index/rstartree"
	"geoindexing_comparison/service/index/rtree"
)

var Indexes = []index.NewIndex{
	kdtree.New,
	rtree.New,
	rstartree.New,
	quadtree.New,
	bruteforce.New,
}

var NameToNewIndex = func() map[string]index.NewIndex {
	mapping := make(map[string]index.NewIndex)
	for _, newIndex := range Indexes {
		mapping[newIndex().Name()] = newIndex
	}
	return mapping
}()
