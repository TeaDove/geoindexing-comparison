package addapter

import (
	"geoindexing_comparison/geo"
)

type Collection interface {
	// Init
	// Initialize empty collection
	// Allowed to be unoptimized
	Init()
	// FromArray creates Collection from geo.Points
	// Allowed to be unoptimized
	FromArray(points geo.Points)
	// Points returns Collection's geo.Points
	// Allowed to be unoptimized
	Points() geo.Points

	// Insert inserts geo.Point to Collection
	Insert(point geo.Point)
	// Remove removes point from geo.Point
	Remove(point geo.Point)

	// KNN returns array of closest n geo.Points to given geo.Point
	KNN(point geo.Point, n int) geo.Points
	// RangeSearch returns all geo.Points in radius around geo.Point
	// RangeSearch(point geo.Point, radius float64) geo.Points

	// String returns string representation of Collection
	// Allowed to be unoptimized
	String() string
}
