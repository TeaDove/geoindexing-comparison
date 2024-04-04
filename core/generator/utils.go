package generator

import (
	"geoindexing_comparison/core/utils"
	"math/rand"
)

func randFloat(min, max float64) float64 {
	return utils.ToFixed(min+rand.Float64()*(max-min), 6)
}
