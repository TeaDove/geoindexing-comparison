package stats

import (
	"fmt"
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

type Array[T constraints.Integer] []T

func NewArray[T constraints.Integer](dur []T) Array[T] {
	sort.Slice(dur, func(i, j int) bool {
		return dur[i] < dur[j]
	})

	return dur
}

func (r Array[T]) String() string {
	return fmt.Sprintf("avg: %f, median: %d, min: %d, max: %d, p90: %d, p95: %d, p99: %d, len: %d",
		r.Avg(),
		r.Median(),
		r.Min(),
		r.Max(),
		r.Quantile(90),
		r.Quantile(95),
		r.Quantile(99),
		len(r),
	)
}

func (r Array[T]) QualifiedAvg() float64 {
	if len(r) < 3 {
		return r.Avg()
	}

	qualifiedArr := r[1 : len(r)-1]
	return qualifiedArr.Avg()
}

func (r Array[T]) Avg() float64 {
	if len(r) == 0 {
		return 0
	}

	var sum T
	for _, el := range r {
		sum += el
	}

	return float64(sum) / float64(len(r))
}

func (r Array[T]) Median() T {
	if len(r) == 0 {
		return 0
	}

	return r[len(r)/2]
}

func (r Array[T]) Max() T {
	if len(r) == 0 {
		return 0
	}

	return r[len(r)-1]
}

func (r Array[T]) Min() T {
	if len(r) == 0 {
		return 0
	}

	return r[0]
}

func (r Array[T]) Quantile(p float64) T {
	if len(r) == 0 {
		return 0
	}

	idx := int(p * float64(len(r)) / 100)

	return r[idx]
}

func (r Array[T]) Variance() float64 {
	var sum T
	for _, dur := range r {
		sum += dur
	}

	mean := float64(sum) / float64(len(r))

	var sumDif float64
	for _, el := range r {
		sumDif += math.Pow(float64(el)-mean, 2)
	}

	return sumDif / float64(len(r))
}

func (r Array[T]) SE() float64 {
	return math.Sqrt(r.Variance())
}

func (r Array[T]) AvgWithSE() [3]float64 {
	if len(r) == 0 {
		return [3]float64{0, 0, 0}
	}

	avg := r.Avg()
	se := r.SE()

	return [3]float64{avg - se, avg, avg + se}
}
