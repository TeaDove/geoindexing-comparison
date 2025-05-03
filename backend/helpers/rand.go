package helpers

import "math/rand/v2"

var (
	Seed = rand.NewPCG(42, 69)
	RNG  = rand.New(Seed)
)
