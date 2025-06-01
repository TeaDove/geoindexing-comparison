package manager_service

import (
	"context"
	"github.com/pkg/errors"
	"slices"
	"time"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Dataset struct {
	Points           []Point `json:"points"`
	RegressionPoints []Point `json:"regressionPoints,omitempty"`
}

func (r *Service) GetChartPoints(ctx context.Context, runID int) (map[string]map[string]Dataset, error) {
	stats, err := r.repository.GetStats(ctx, runID)
	if err != nil {
		return nil, errors.Wrap(err, "could not get stats")
	}

	taskToIndex := make(map[string]map[string]Dataset)

	for _, stat := range stats {
		_, ok := taskToIndex[stat.Task]
		if !ok {
			taskToIndex[stat.Task] = map[string]Dataset{}
		}

		dataset, _ := taskToIndex[stat.Task][stat.Index]

		dataset.Points = append(dataset.Points, Point{
			X: float64(stat.Amount),
			Y: stat.Durs.QualifiedAvg() / float64(time.Microsecond),
		})

		slices.SortFunc(dataset.Points, func(a, b Point) int {
			if a.X > b.X {
				return 1
			}
			return -1
		})

		dataset.RegressionPoints = statsLinreg(dataset.Points)

		taskToIndex[stat.Task][stat.Index] = dataset
	}

	return taskToIndex, nil
}

func statsLinreg(points []Point) []Point {
	if len(points) <= 2 {
		return nil
	}

	a, b := linreg(points)
	// Y = aX + b
	return []Point{
		{
			X: points[0].X,
			Y: a*points[0].X + b,
		},
		{
			X: points[len(points)-1].X,
			Y: a*points[len(points)-1].X + b,
		},
	}
}

func linreg(points []Point) (float64, float64) {
	n := float64(len(points))
	var sx, sy, sxx, syy, sxy float64
	for _, point := range points {
		sx += point.X
		sy += point.Y
		sxx += point.X * point.X
		syy += point.Y * point.Y
		sxy += point.X * point.Y
	}
	det := sxx*n - sx*sx
	return (sxy*n - sy*sx) / det, (sxx*sy - sx*sxy) / det
}
