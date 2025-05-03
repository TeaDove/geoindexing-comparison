package helpers

import (
	"strings"
)

func ID() string {
	const (
		alfabet = "0123456789abcdefghijklmnopqrstuvwxyz"
		dash    = "-"
	)

	var builder strings.Builder
	for range 4 {
		builder.WriteByte(alfabet[RNG.IntN(len(alfabet))])
	}
	builder.WriteString(dash)
	for range 4 {
		builder.WriteByte(alfabet[RNG.IntN(len(alfabet))])
	}

	return builder.String()
}
