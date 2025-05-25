package helpers

import (
	"math/rand/v2"
	"strings"
)

func ID(rng *rand.Rand) string {
	const (
		alfabet = "0123456789abcdefghijklmnopqrstuvwxyz"
		dash    = "-"
	)

	var builder strings.Builder
	for range 4 {
		builder.WriteByte(alfabet[rng.IntN(len(alfabet))])
	}

	builder.WriteString(dash)

	for range 4 {
		builder.WriteByte(alfabet[rng.IntN(len(alfabet))])
	}

	return builder.String()
}
