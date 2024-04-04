package main

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases"
	"geoindexing_comparison/core/cases/tasks"
)

func main() {
	allCases := AllCases()
	cases.Run(&allCases)
}

func AllCases() cases.RunCase {
	return cases.RunCase{
		Collections: addapter_all.All(),
		Tasks:       tasks.All(),
		AmountStart: 10_000,
		AmountEnd:   100_000,
		AmountStep:  1000,
	}
}
