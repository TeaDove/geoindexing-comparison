package generator

import (
	"math/rand"

	"geoindexing_comparison/core/utils"
)

func randFloat(min, max float64) float64 {
	return utils.ToFixed(min+rand.Float64()*(max-min), 6)
}
