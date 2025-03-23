package index

import (
	"geoindexing_comparison/service/geo"
	"time"
)

type NewIndex func() IndexImpl

type IndexImpl interface {
	// FromArray creates IndexImpl from geo.Points
	// Allowed to be unoptimized
	FromArray(points geo.Points)

	// InsertTimed inserts geo.Point to IndexImpl
	InsertTimed(point geo.Point) time.Duration

	// KNNTimed returns array of closest n geo.Points to given geo.Point
	KNNTimed(point geo.Point, n uint64) (geo.Points, time.Duration)

	// RangeSearchTimed returns run geo.Points in radius around geo.Point
	RangeSearchTimed(point geo.Point, radius float64) (geo.Points, time.Duration)

	// String returns string representation of IndexImpl
	// Allowed to be unoptimized
	String() string
}

type IndexInfo struct {
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
}

type Index struct {
	Builder NewIndex  `json:"-"`
	Info    IndexInfo `json:"info"`
}
