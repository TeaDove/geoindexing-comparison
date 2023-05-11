package generator

import (
	"math/rand"
)

const (
	latLowerBoundDefault float64 = 48.226506
	latUpperBoundDefault float64 = 58.016099
	lonLowerBoundDefault float64 = 27.178174
	lonUpperBoundDefault float64 = 41.563808
	pointsAmountDefault  int     = 5_000
)

type Generator struct {
	LatLowerBound float64
	LatUpperBound float64
	LonLowerBound float64
	LonUpperBound float64
	PointsAmount  int
}

var DefaultGenerator = GetDefaultGenerator()

func GetDefaultGenerator() Generator {
	return Generator{
		LatLowerBound: latLowerBoundDefault,
		LatUpperBound: latUpperBoundDefault,
		LonLowerBound: lonLowerBoundDefault,
		LonUpperBound: lonUpperBoundDefault,
		PointsAmount:  pointsAmountDefault,
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
