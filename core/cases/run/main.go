package main

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases"
	"geoindexing_comparison/core/cases/tasks"
)

func main() {
	cases.Run(&allCasesFast)
}

var (
	allCases = cases.RunCase{
		Collections: addapter_all.All(),
		Tasks:       tasks.All(),
		AmountStart: 10_000,
		AmountEnd:   100_000,
		AmountStep:  1000,
	}

	allCasesSmallAmount = cases.RunCase{
		Collections: addapter_all.All(),
		Tasks:       tasks.All(),
		AmountStart: 100,
		AmountEnd:   1000,
		AmountStep:  10,
	}

	allCasesFast = cases.RunCase{
		Collections: addapter_all.All(),
		Tasks:       tasks.All(),
		AmountStart: 100,
		AmountEnd:   1000,
		AmountStep:  100,
	}
)
