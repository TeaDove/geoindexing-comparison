package generator

import (
	"math/rand"

	"github.com/teadove/teasutils/utils/conv_utils"
)

func randFloat(bottom, top float64) float64 {
	return conv_utils.ToFixed(bottom+rand.Float64()*(top-bottom), 6)
}
