package main

import (
	"geoindexing_comparison/cases"
	"geoindexing_comparison/cases/tasks"
)

func main() {
	cases.Run(AllCases()...)
}

func AllCases() []cases.RunCase {
	runCases := make([]cases.RunCase, 0, 10)
	for _, collection := range cases.AllCollections() {
		for _, task := range tasks.AllTasks() {
			runCases = append(runCases, cases.RunCase{
				Collection:  collection,
				Task:        task,
				Repetitions: 100,
			})
		}
	}

	return runCases
}
