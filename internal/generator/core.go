package generator

import (
	"math/rand"
)

type Generator struct {
	LatLowerBound    float64
	LatUpperBound    float64
	LonLowerBound    float64
	LonUpperBound    float64
	PointsAmount     int
	KNNSearchSize    int
	RadiusSearchSize float64
}

var DefaultGenerator = NewDefaultGenerator()

func NewDefaultGenerator() Generator {
	return Generator{
		LatLowerBound:    48.226506,
		LatUpperBound:    58.016099,
		LonLowerBound:    27.178174,
		LonUpperBound:    41.563808,
		PointsAmount:     1_000,
		KNNSearchSize:    10,
		RadiusSearchSize: 0.5,
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
