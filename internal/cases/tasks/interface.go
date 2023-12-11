package cases

import (
	"geoindexing_comparison/addapter"
	"time"
)

type Task interface {
	Name() string
	Run(col addapter.Collection) []time.Duration
}
