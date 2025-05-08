package geo

import (
	"github.com/mmcloughlin/geohash"
	"iter"
)

func geohashNextDirection(direction geohash.Direction) geohash.Direction {
	switch direction {
	case geohash.North:
		return geohash.East
	case geohash.East:
		return geohash.South
	case geohash.South:
		return geohash.West
	case geohash.West:
		return geohash.North
	}
	panic("unreachable")
}

func GeohashNeighborIter(v uint64, bits uint) iter.Seq[uint64] {
	return func(yield func(uint64) bool) {
		var (
			step      = 1
			direction = geohash.North
		)

		for {
			for range step {
				v = geohash.NeighborIntWithPrecision(v, bits, direction)
				if !yield(v) {
					return
				}
			}

			direction = geohashNextDirection(direction)
			for range step {
				v = geohash.NeighborIntWithPrecision(v, bits, direction)
				if !yield(v) {
					return
				}
			}

			direction = geohashNextDirection(direction)
			step++
		}
	}
}

func GeohashCircled(lat, lng float64, radius float64, bits uint) []uint64 {
	return nil
}
