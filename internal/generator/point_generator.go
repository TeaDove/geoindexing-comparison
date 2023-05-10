package generator

func (r *Generator) GeneratePoint() Point {
	coord := Point{
		Lat: r.randLat(),
		Lon: r.randLon(),
	}
	return coord
}

func (r *Generator) GeneratePoints(amount int) []Point {
	var points = make([]Point, amount)
	for i := 0; i < amount; i++ {
		points[i] = r.GeneratePoint()
	}
	return points
}
