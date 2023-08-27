package geo

func (r Point) Dimensions() int {
	return 2
}

func (r Point) Dimension(i int) float64 {
	switch i {
	case 0:
		return r.Lat
	default:
		return r.Lon
	}
}
