package addapter

import "geoindexing_comparison/geo"

type Collection interface {
	// KNN returns array of closest n geo.Points to given geo.Point
	KNN(point geo.Point, n int) geo.Points
	// RangeSearch returns all geo.Points in radius around geo.Point
	RangeSearch(point geo.Point, radius float64) geo.Points

	// Empty creates empty Collection
	Empty()

	// FromArray creates Collection from geo.Points
	FromArray(points geo.Points)
	// Points returns Collection's geo.Points
	Points() geo.Points
	// Insert inserts geo.Point to Collection
	Insert(point geo.Point)
	// Remove removes point from geo.Point
	Remove(point geo.Point)

	// String returns string representation of Collection
	String() string
}
