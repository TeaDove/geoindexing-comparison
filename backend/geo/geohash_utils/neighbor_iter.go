package geohash_utils

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

func NeighborIter(v uint64, bits uint) iter.Seq[uint64] {
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

func NeighborIterSquared(v uint64, bits uint) iter.Seq[[]uint64] {
	return func(yield func([]uint64) bool) {
		if !yield([]uint64{v}) {
			return
		}

		bottomLeft := geohash.NeighborIntWithPrecision(v, bits, geohash.SouthWest)

		step := 2
		for {
			perimeter := collectPerimeter(bottomLeft, bits, step, step)
			if !yield(perimeter) {
				return
			}
			step += 2
			bottomLeft = geohash.NeighborIntWithPrecision(bottomLeft, bits, geohash.SouthWest)
		}
	}
}
