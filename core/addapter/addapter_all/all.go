package addapter_all

import (
	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/addapter/kdtree"
	"geoindexing_comparison/core/addapter/rtree"
)

type CollectionInit func() addapter.Collection

func All() []CollectionInit {
	return []CollectionInit{
		rtree.New,
		kdtree.New,
	}
}
