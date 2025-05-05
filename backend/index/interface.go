package index

import (
	"geoindexing_comparison/backend/geo"
	"time"
)

type NewIndex func() Impl

type Impl interface {
	// FromArray creates Impl from geo.Points
	// Allowed to be unoptimized
	FromArray(points geo.Points)

	ToArray() geo.Points

	// String returns string representation of Impl
	// Allowed to be unoptimized
	String() string

	// InsertTimed inserts geo.Point to Impl
	InsertTimed(point geo.Point) time.Duration

	// KNNTimed returns array of closest n geo.Points to given geo.Point
	KNNTimed(origin geo.Point, n uint64) (geo.Points, time.Duration)

	// RangeSearchTimed returns run geo.Points in radius around geo.Point
	RangeSearchTimed(origin geo.Point, radius float64) (geo.Points, time.Duration)
}

type Info struct {
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
}

type Index struct {
	Builder NewIndex `json:"-"`
	Info    Info     `json:"info"`
}
