package addapter

import "geoindexing_comparison/geo"

type Collection interface {
	KNN(point geo.Point, n int) []geo.Point
	RangeSearch(radius float64) []geo.Point

	Empty()
	FromArray(points []geo.Point)
	Insert(point geo.Point)
	Remove(point geo.Point)
	String() string
}
