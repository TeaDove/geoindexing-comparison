package main

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases"
	"geoindexing_comparison/core/cases/tasks"
)

func init() {
	//utils.Init()
}

func main() {
	cases.Run(AllCases()...)
}

func AllCases() []cases.RunCase {
	runCases := make([]cases.RunCase, 0, 10)
	for _, collection := range addapter_all.All() {
		for _, task := range tasks.All() {
			runCases = append(runCases, cases.RunCase{
				Collection:  collection,
				Task:        task,
				AmountStart: 100,
				AmountEnd:   1_000,
				AmountStep:  100,
			})
		}
	}

	return runCases
}
