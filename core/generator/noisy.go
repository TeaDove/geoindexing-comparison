package generator

import (
	"geoindexing_comparison/core/geo"
	"geoindexing_comparison/core/utils"
	"github.com/KEINOS/go-noise"
	"time"
)

type NoisyGenerator struct {
}

func NewNoisyGenerator() *NoisyGenerator {
	return &NoisyGenerator{}
}

func (r *NoisyGenerator) Points(amount int) geo.Points {
	_, err := noise.New(noise.Perlin, time.Now().Unix())
	utils.Check(err)

	return nil
}
