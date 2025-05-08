package geo

import (
	"github.com/mmcloughlin/geohash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeohashNeighborIter(t *testing.T) {
	t.Parallel()

	chars := uint(6)
	bits := chars * 5

	lat, lng := geohash.Decode("ucfv7t")
	originGeohash := geohash.EncodeIntWithPrecision(lat, lng, bits)
	idx := 0
	exp := []string{"ucfv7w", "ucfv7y", "ucfv7v", "ucfv7u", "ucfv7s", "ucfv7k", "ucfv7m", "ucfv7q", "ucfv7r", "ucfv7x"}

	for neighbor := range GeohashNeighborIter(originGeohash, bits) {
		lat, lng = geohash.DecodeIntWithPrecision(neighbor, bits)
		hash := geohash.EncodeWithPrecision(lat, lng, chars)
		assert.Equal(t, exp[idx], hash)

		idx++
		if idx >= 10 {
			break
		}
	}
}
