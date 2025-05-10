package geohash_utils

import (
	"fmt"
	"geoindexing_comparison/backend/geo"
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

	for neighbor := range NeighborIter(originGeohash, bits) {
		lat, lng = geohash.DecodeIntWithPrecision(neighbor, bits)
		hash := geohash.EncodeWithPrecision(lat, lng, chars)
		assert.Equal(t, exp[idx], hash)

		idx++
		if idx >= 10 {
			break
		}
	}
}

func TestGeohashNeighborIterSquared(t *testing.T) {
	t.Parallel()

	chars := uint(6)
	bits := chars * 5

	lat, lng := geohash.Decode("ucfv7t")
	originGeohash := geohash.EncodeIntWithPrecision(lat, lng, bits)
	idx := 0
	exp := [][]string{{"ucfv7t"}, {"ucfv7k", "ucfv7m", "ucfv7q", "ucfv7w", "ucfv7y", "ucfv7v", "ucfv7u", "ucfv7s"}, {"ucfv75", "ucfv7h", "ucfv7j", "ucfv7n", "ucfv7p", "ucfv7r", "ucfv7x", "ucfv7z", "ucfvkp", "ucfvkn", "ucfvkj", "ucfvkh", "ucfvk5", "ucfv7g", "ucfv7e", "ucfv77"}}

	for neighbors := range NeighborIterSquared(originGeohash, bits) {
		for jdx, neighbor := range neighbors {
			lat, lng = geohash.DecodeIntWithPrecision(neighbor, bits)
			hash := geohash.EncodeWithPrecision(lat, lng, chars)
			assert.Equal(t, exp[idx][jdx], hash)
		}

		idx++
		if idx >= 2 {
			break
		}
	}
}

func TestGeohashBboxPerimeterOk(t *testing.T) {
	t.Parallel()

	chars := uint(6)
	bits := chars * 5
	bottomLeft := geo.NewPoint(55.793217, 37.781136)
	upperRight := geo.NewPoint(55.801616, 37.803966)

	bbox := NewBBox(bottomLeft.Lat, bottomLeft.Lon, upperRight.Lat, upperRight.Lon, bits)
	expHashed := []string{"ucfv7s", "ucfv7t", "ucfv7w", "ucfv7y", "ucfvkn", "ucfvkq", "ucfvkm", "ucfvkk", "ucfvkh", "ucfv7u"}

	for idx, hash := range bbox.Perimeter() {
		lat, lng := geohash.DecodeIntWithPrecision(hash, bits)
		assert.Equal(t, expHashed[idx], geohash.EncodeWithPrecision(lat, lng, chars))
	}
}

func TestGeohashBboxInnterOk(t *testing.T) {
	t.Parallel()

	chars := uint(6)
	bits := chars * 5
	bottomLeft := geo.NewPoint(55.793217, 37.781136)
	upperRight := geo.NewPoint(55.801616, 37.803966)

	bbox := NewBBox(bottomLeft.Lat, bottomLeft.Lon, upperRight.Lat, upperRight.Lon, bits)
	expHashed := []string{"ucfv7v", "ucfvkj"}

	for idx, hash := range bbox.Inner() {
		lat, lng := geohash.DecodeIntWithPrecision(hash, bits)
		assert.Equal(t, expHashed[idx], geohash.EncodeWithPrecision(lat, lng, chars))
	}
}

func TestGeohashFastNeighborsOk(t *testing.T) {
	chars := uint(6)
	bits := chars * 5

	lat, lng := geohash.Decode("ucfv7t")
	originGeohash := geohash.EncodeIntWithPrecision(lat, lng, bits)

	directions := []geohash.Direction{geohash.North, geohash.NorthEast, geohash.East, geohash.SouthEast, geohash.South, geohash.SouthWest, geohash.West, geohash.NorthWest}
	for _, direction := range directions {
		t.Run(fmt.Sprint(direction), func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, NeighborIntWithPrecision(originGeohash, bits, direction), geohash.NeighborIntWithPrecision(originGeohash, bits, direction))
		})
	}
}
