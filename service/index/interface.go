package index

import (
	"geoindexing_comparison/service/geo"
	"time"
)

type NewIndex func() Index

type Index interface {
	// Name returns name of struct
	// Allowed to be unoptimized
	Name() string

	// FromArray creates Index from geo.Points
	// Allowed to be unoptimized
	FromArray(points geo.Points)

	// InsertTimed inserts geo.Point to Index
	InsertTimed(point geo.Point) time.Duration

	// KNNTimed returns array of closest n geo.Points to given geo.Point
	KNNTimed(point geo.Point, n int) (geo.Points, time.Duration)

	// RangeSearchTimed returns run geo.Points in radius around geo.Point
	RangeSearchTimed(point geo.Point, radius float64) (geo.Points, time.Duration)

	// String returns string representation of Index
	// Allowed to be unoptimized
	String() string
}
