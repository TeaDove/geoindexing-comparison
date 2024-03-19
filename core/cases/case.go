package cases

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases/tasks"
)

type RunCase struct {
	Collection  addapter_all.CollectionInit
	Task        tasks.Task
	AmountStart int
	AmountEnd   int
	AmountStep  int
}
