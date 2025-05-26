package helpers

import "math/rand/v2"

const (
	constSeed1 uint64 = 42
	constSeed2 uint64 = 69
)

func RNG() *rand.Rand {
	return NewRNG(constSeed1, constSeed2)
}

func NewRNG(seed1, seed2 uint64) *rand.Rand {
	return rand.New(rand.NewPCG(seed1, seed2))
}
