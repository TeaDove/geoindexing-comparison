package main

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases"
	"geoindexing_comparison/core/cases/tasks"
)

func main() {
	results := cases.Run(AllCases())
	drawResults(results)
}

func AllCases() []cases.RunCase {
	runCases := make([]cases.RunCase, 0, 10)
	for _, collection := range addapter_all.All() {
		for _, task := range tasks.All() {
			runCases = append(runCases, cases.RunCase{
				Collection:  collection,
				Task:        task,
				AmountStart: 100,
				AmountEnd:   1000,
				AmountStep:  10,
			})
		}
	}

	return runCases
}
