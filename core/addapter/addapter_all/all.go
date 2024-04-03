package addapter_all

import (
	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/addapter/bruteforce"
	"geoindexing_comparison/core/addapter/kdtree"
	"geoindexing_comparison/core/addapter/quadtree"
	"geoindexing_comparison/core/addapter/rstartree"
	"geoindexing_comparison/core/addapter/rtree"
)

type CollectionInit func() addapter.Collection

func All() []CollectionInit {
	return []CollectionInit{
		kdtree.New,
		rtree.New,
		rstartree.New,
		quadtree.New,
		bruteforce.New,
	}
}
