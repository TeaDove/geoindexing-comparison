package generator

import (
	"math/rand"
)

const (
	latLowerBoundDefault float64 = 0
	latUpperBoundDefault float64 = 180
	lonLowerBoundDefault float64 = -90
	lonUpperBoundDefault float64 = 90
)

type Generator struct {
	LatLowerBound float64
	LatUpperBound float64
	LonLowerBound float64
	LonUpperBound float64
}

func DefaultGenerator() Generator {
	return Generator{
		LatLowerBound: latLowerBoundDefault,
		LatUpperBound: latUpperBoundDefault,
		LonLowerBound: lonLowerBoundDefault,
		LonUpperBound: lonUpperBoundDefault,
	}
}

func (r *Generator) randLat() float64 {
	return randFloat(r.LatLowerBound, r.LatUpperBound)
}

func (r *Generator) randLon() float64 {
	return randFloat(r.LonLowerBound, r.LonUpperBound)
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = randFloat(min, max)
	}
	return res
}
