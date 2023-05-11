package addapter

import "geoindexing_comparison/geo"

type Collection interface {
	KNN(point geo.Point, n int) geo.Points
	RangeSearch(radius float64) geo.Points

	Empty()
	FromArray(points geo.Points)
	Points() geo.Points
	Insert(point geo.Point)
	Remove(point geo.Point)
	String() string
}

//func PaintInCollection(r *Collection, points geo.Points, category geo.Category) {
//	for _
//}
