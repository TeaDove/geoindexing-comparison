// Package geohash_utils
//
// Geohash utils, such as BBox
package geohash_utils

import (
	"github.com/mmcloughlin/geohash"
)

type BBox struct {
	bits       uint
	leftBottom uint64

	height int
	wight  int
}

func NewBBox(bottomLeftLat, bottomLeftLon, upperRightLat, upperRightLon float64, bits uint) BBox {
	var (
		wight  = 0
		height = 0

		bottomLeftHash  = geohash.EncodeIntWithPrecision(bottomLeftLat, bottomLeftLon, bits)
		upperLeftHash   = geohash.EncodeIntWithPrecision(upperRightLat, bottomLeftLon, bits)
		bottomRightHash = geohash.EncodeIntWithPrecision(bottomLeftLat, upperRightLon, bits)
	)

	for bottomLeftHash != upperLeftHash {
		upperLeftHash = NeighborIntWithPrecision(upperLeftHash, bits, geohash.South)
		height++
	}

	for bottomLeftHash != bottomRightHash {
		bottomRightHash = NeighborIntWithPrecision(bottomRightHash, bits, geohash.West)
		wight++
	}

	r := BBox{
		bits:       bits,
		leftBottom: bottomLeftHash,
		height:     height,
		wight:      wight,
	}

	return r
}

func collectPerimeter(hash uint64, bits uint, height, wight int) []uint64 {
	perimeter := make([]uint64, 0, height)

	for range height {
		perimeter = append(perimeter, hash)
		hash = NeighborIntWithPrecision(hash, bits, geohash.North)
	}

	for range wight {
		perimeter = append(perimeter, hash)
		hash = NeighborIntWithPrecision(hash, bits, geohash.East)
	}

	for range height {
		perimeter = append(perimeter, hash)
		hash = NeighborIntWithPrecision(hash, bits, geohash.South)
	}

	for range wight {
		perimeter = append(perimeter, hash)
		hash = NeighborIntWithPrecision(hash, bits, geohash.West)
	}

	return perimeter
}

// Perimeter
// Returns outer part of BBox.
func (r *BBox) Perimeter() []uint64 {
	return collectPerimeter(r.leftBottom, r.bits, r.height, r.wight)
}

// Inner
// Returns inner part of BBox, points from Perimeter are not included.
func (r *BBox) Inner() []uint64 {
	var (
		height    = r.height - 1
		wight     = r.wight - 1
		inner     = make([]uint64, 0, r.height)
		hash      = NeighborIntWithPrecision(r.leftBottom, r.bits, geohash.NorthEast)
		innerHash = hash //nolint: ineffassign, wastedassign // have no idea how to do
	)

	for range height {
		innerHash = hash
		for range wight {
			inner = append(inner, innerHash)
			innerHash = NeighborIntWithPrecision(innerHash, r.bits, geohash.East)
		}

		hash = NeighborIntWithPrecision(hash, r.bits, geohash.North)
	}

	return inner
}
