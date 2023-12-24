package addapter_all

import (
	"geoindexing_comparison/addapter"
	"geoindexing_comparison/addapter/kdtree"
	"geoindexing_comparison/addapter/rtree"
)

type CollectionInit func() addapter.Collection

func All() []CollectionInit {
	return []CollectionInit{
		rtree.New,
		kdtree.New,
	}
}
