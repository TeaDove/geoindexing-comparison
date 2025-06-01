package manager_repository

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/teadove/teasutils/utils/test_utils"
	"golang.org/x/exp/slices"
	"testing"
)

/*
def linreg(X: list[float], Y: list[float]) -> tuple[float, float]:
    """
    return a,b in solution to y = ax + b such that root mean square distance between trend line and original points is minimized
    """
    N = len(X)
    Sx = Sy = Sxx = Syy = Sxy = 0.0
    for x, y in zip(X, Y):
        Sx = Sx + x
        Sy = Sy + y
        Sxx = Sxx + x*x
        Syy = Syy + y*y
        Sxy = Sxy + x*y
    det = Sxx * N - Sx * Sx
    return (Sxy * N - Sy * Sx)/det, (Sxx * Sy - Sx * Sxy)/det
*/

func linreg(xs []float64, ys []float64) (float64, float64) {
	if len(xs) != len(ys) {
		panic(errors.New("len of X and Y must be equal"))
	}

	n := float64(len(xs))
	var sx, sy, sxx, syy, sxy float64
	for idx := range len(xs) {
		x := xs[idx]
		y := ys[idx]

		sx += x
		sy += y
		sxx += x * x
		syy += y * y
		sxy += x * y
	}
	det := sxx*n - sx*sx
	return (sxy*n - sy*sx) / det, (sxx*sy - sx*sxy) / det
}

func TestSpeedOk(t *testing.T) {
	t.Parallel()

	runID, index, task := 40, "rtree", "knn_25_p"
	ctx := test_utils.GetLoggedContext()
	r, err := NewRepository(ctx, "/Users/pibragimov/projects/geoindexing-comparison/.data/db.sqlite")
	require.NoError(t, err)

	stats, err := r.GetStats(ctx, runID)
	require.NoError(t, err)
	var testStats []Stats
	for _, stat := range stats {
		if stat.Index == index && stat.Task == task {
			testStats = append(testStats, stat)
		}
	}

	slices.SortFunc(testStats, func(a, b Stats) int {
		if a.Idx > b.Idx {
			return 1
		}
		return -1
	})

	for _, stat := range testStats {
		test_utils.Pprint(stat.Durs.Avg())
	}

	var speed []float64
	for idx, stat := range testStats {
		if idx+1 >= len(testStats) {
			break
		}

		speed = append(speed, testStats[idx+1].Durs.QualifiedAvg()-stat.Durs.QualifiedAvg())
	}

	test_utils.Pprint(speed)

	var xs, ys []float64
	for _, stat := range testStats {
		xs = append(xs, float64(stat.Amount))
		ys = append(ys, stat.Durs.QualifiedAvg())
	}

	test_utils.Pprint(linreg(xs, ys))
}
