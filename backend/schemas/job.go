package schemas

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/helpers/stats"
	"time"
)

type Job struct {
	Task  string `json:"task"`
	Index string `json:"index"`

	Amount      int        `json:"amount"`
	Points      geo.Points `json:"points"`
	RandomPoint geo.Point  `json:"randomPoint" json:"randomPoint"`
}

type JobResult struct {
	Index  string
	Task   string
	Amount int

	Durs stats.Array[time.Duration]
}
