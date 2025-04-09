package generator

import (
	"geoindexing_comparison/backend/geo"

	"time"

	"github.com/KEINOS/go-noise"
)

type NoisyGenerator struct{}

func NewNoisyGenerator() *NoisyGenerator {
	return &NoisyGenerator{}
}

func (r *NoisyGenerator) Points(amount int) geo.Points {
	_, err := noise.New(noise.Perlin, time.Now().Unix())
	if err != nil {
		panic(err)
	}

	return nil
}
