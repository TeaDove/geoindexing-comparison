package main

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases"
	"geoindexing_comparison/core/cases/tasks"
)

func main() {
	cases.Run(&allCasesMedium)
}

var (
	allCases = cases.RunCase{
		Name:        "big",
		Collections: addapter_all.All(),
		Tasks:       tasks.KnnAndRadiusSearch,
		AmountStart: 10_000,
		AmountEnd:   100_000,
		AmountStep:  1000,
	}

	allCasesMedium = cases.RunCase{
		Name:        "medium",
		Collections: addapter_all.All(),
		Tasks:       tasks.OnlyRadiusSearch,
		AmountStart: 10_000,
		AmountEnd:   100_000,
		AmountStep:  3_000,
	}

	allCasesSmallAmount = cases.RunCase{
		Name:        "small_amount",
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
