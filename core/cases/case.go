package cases

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"geoindexing_comparison/core/cases/tasks"
)

type RunCase struct {
	Name        string
	Collections []addapter_all.CollectionInit
	Tasks       []tasks.Task
	AmountStart int
	AmountEnd   int
	AmountStep  int
}
