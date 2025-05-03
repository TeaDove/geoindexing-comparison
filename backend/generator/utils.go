package generator

import (
	"geoindexing_comparison/backend/helpers"
	"github.com/teadove/teasutils/utils/conv_utils"
)

func randFloat(bottom, top float64) float64 {
	return conv_utils.ToFixed(bottom+helpers.RNG.Float64()*(top-bottom), 6) //nolint: gosec // Allowed here
}
