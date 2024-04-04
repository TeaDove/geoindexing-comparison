package main

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases"
	"geoindexing_comparison/core/cases/tasks"
)

func main() {
	allCases := AllCases()
	results := cases.Run(&allCases)
	drawResults(results)
}

func AllCases() cases.RunCase {
	return cases.RunCase{
		Collections: addapter_all.All(),
		Tasks:       tasks.All()[:1],
		AmountStart: 100,
		AmountEnd:   1000,
		AmountStep:  100,
	}
}
