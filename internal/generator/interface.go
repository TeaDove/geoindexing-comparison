package generator

import "geoindexing_comparison/geo"

type Generator interface {
	Point(input *Input) geo.Point
	Points(input *Input, amount int) geo.Points
}
