package cases

import (
	"geoindexing_comparison/addapter"
	"geoindexing_comparison/addapter/kdtree"
	"geoindexing_comparison/addapter/rtree"
	"geoindexing_comparison/generator"
)

type RunCase struct {
	Collection  func() addapter.Collection
	Task        Task
	Repetitions int
	Generator   generator.Generator
}

func AllCollections() []func() addapter.Collection {
	return []func() addapter.Collection{
		rtree.New,
		kdtree.New,
	}
}
