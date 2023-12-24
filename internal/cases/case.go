package cases

import (
	"geoindexing_comparison/addapter/addapter_all"
	"geoindexing_comparison/cases/tasks"
)

type RunCase struct {
	Collection  addapter_all.CollectionInit
	Task        tasks.Task
	AmountStart int
	AmountEnd   int
	AmountStep  int
}
