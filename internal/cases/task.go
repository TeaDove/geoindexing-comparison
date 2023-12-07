package cases

import (
	"geoindexing_comparison/addapter"
	"geoindexing_comparison/geo"
	"time"
)

type Task interface {
	Name() string
	Run(col addapter.Collection, points geo.Points) time.Duration
}
