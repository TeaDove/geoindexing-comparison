package cases

import (
	"geoindexing_comparison/service/cases/tasks"
	"geoindexing_comparison/service/index"
)

type RunCase struct {
	Name        string
	Indexes     []index.NewIndex
	Tasks       []tasks.Task
	AmountStart int
	AmountEnd   int
	AmountStep  int
}
