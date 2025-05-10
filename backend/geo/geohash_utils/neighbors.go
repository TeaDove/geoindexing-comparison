package geohash_utils

import "github.com/mmcloughlin/geohash"

func NeighborIntWithPrecision(hash uint64, bits uint, direction geohash.Direction) uint64 {
	box := geohash.BoundingBoxIntWithPrecision(hash, bits)
	lat, lng := box.Center()
	latDelta := box.MaxLat - box.MinLat
	lngDelta := box.MaxLng - box.MinLng

	switch direction {
	case geohash.North:
		return geohash.EncodeIntWithPrecision(lat+latDelta, lng, bits)
	case geohash.NorthEast:
		return geohash.EncodeIntWithPrecision(lat+latDelta, lng+lngDelta, bits)
	case geohash.East:
		return geohash.EncodeIntWithPrecision(lat, lng+lngDelta, bits)
	case geohash.SouthEast:
		return geohash.EncodeIntWithPrecision(lat-latDelta, lng+lngDelta, bits)
	case geohash.South:
		return geohash.EncodeIntWithPrecision(lat-latDelta, lng, bits)
	case geohash.SouthWest:
		return geohash.EncodeIntWithPrecision(lat-latDelta, lng-lngDelta, bits)
	case geohash.West:
		return geohash.EncodeIntWithPrecision(lat, lng-lngDelta, bits)
	case geohash.NorthWest:
		return geohash.EncodeIntWithPrecision(lat+latDelta, lng-lngDelta, bits)
	}

	panic("unreachable")
}
