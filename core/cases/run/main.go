package main

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/addapter/kdtree"
	"geoindexing_comparison/core/addapter/rtree"
	"geoindexing_comparison/core/cases"
	"geoindexing_comparison/core/cases/tasks"
)

func main() {
	cases.Run(&allCasesFast)
}

var (
	allCases = cases.RunCase{
		Name:        "big_no_brute",
		Collections: addapter_all.AllWithoutBruteforce(),
		Tasks:       tasks.KnnAndRadiusSearch,
		AmountStart: 100_000,
		AmountEnd:   1_000_000,
		AmountStep:  30_000,
	}

	kdVsRtree = cases.RunCase{
		Name:        "kd_vs_rtree",
		Collections: []addapter_all.CollectionInit{rtree.New, kdtree.New},
		Tasks:       tasks.OnlyKNN1,
		AmountStart: 100,
		AmountEnd:   60_000,
		AmountStep:  300,
	}

	allCasesMedium = cases.RunCase{
		Name:        "medium",
		Collections: addapter_all.All(),
		Tasks:       tasks.KnnAndRadiusSearch,
		AmountStart: 10_000,
		AmountEnd:   100_000,
		AmountStep:  1_000,
	}

	allCasesSmallAmount = cases.RunCase{
		Name:        "cluster_small_amount",
		Collections: addapter_all.All(),
		Tasks:       tasks.All,
		AmountStart: 100,
		AmountEnd:   1000,
		AmountStep:  10,
	}

	allCasesFast = cases.RunCase{
		Name:        "fast",
		Collections: addapter_all.All(),
		Tasks:       tasks.All,
		AmountStart: 100,
		AmountEnd:   1000,
		AmountStep:  100,
	}
)
