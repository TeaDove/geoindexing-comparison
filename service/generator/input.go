package generator

type Input struct {
	LatLowerBound float64
	LatUpperBound float64
	LonLowerBound float64
	LonUpperBound float64
}

var DefaultInput = Input{
	LatLowerBound: 55.466488,
	LatUpperBound: 55.945060,

	LonLowerBound: 37.231972,
	LonUpperBound: 37.848366,
}
