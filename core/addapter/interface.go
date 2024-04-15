package addapter

import (
	"time"

	"geoindexing_comparison/core/geo"
)

type Collection interface {
	// Name returns name of struct
	// Allowed to be unoptimized
	Name() string

	// FromArray creates Collection from geo.Points
	// Allowed to be unoptimized
	FromArray(points geo.Points)

	// InsertTimed inserts geo.Point to Collection
	InsertTimed(point geo.Point) time.Duration

	// KNNTimed returns array of closest n geo.Points to given geo.Point
	KNNTimed(point geo.Point, n int) (geo.Points, time.Duration)

	// RangeSearchTimed returns run geo.Points in radius around geo.Point
	RangeSearchTimed(point geo.Point, radius float64) (geo.Points, time.Duration)

	// String returns string representation of Collection
	// Allowed to be unoptimized
	String() string
}
