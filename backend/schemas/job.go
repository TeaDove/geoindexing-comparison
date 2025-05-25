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
	RandomPoint geo.Point  `json:"randomPoint"`
}

type JobResult struct {
	Index  string `json:"index"`
	Task   string `json:"task"`
	Amount int    `json:"amount"`

	Durs stats.Array[time.Duration] `json:"durs"`
}
