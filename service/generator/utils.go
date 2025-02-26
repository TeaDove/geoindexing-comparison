package generator

import (
	"github.com/teadove/teasutils/utils/conv_utils"
	"math/rand"
)

func randFloat(bottom, top float64) float64 {
	return conv_utils.ToFixed(bottom+rand.Float64()*(top-bottom), 6)
}
