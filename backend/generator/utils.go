package generator

import (
	"github.com/teadove/teasutils/utils/conv_utils"
	"math/rand/v2"
)

func randFloat(rng *rand.Rand, bottom, top float64) float64 {
	return conv_utils.ToFixed(bottom+rng.Float64()*(top-bottom), 6) //nolint: gosec // Allowed here
}
