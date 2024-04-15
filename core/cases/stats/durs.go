package stats

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Durs []time.Duration

func NewDurs(dur []time.Duration) Durs {
	sort.Slice(dur, func(i, j int) bool {
		return dur[i] < dur[j]
	})

	return dur
}

func (r Durs) String() string {
	return fmt.Sprintf("avg: %s, median: %s, min: %s, max: %s, p90: %s, p95: %s, p99: %s",
		r.Avg().String(),
		r.Median().String(),
		r.Min().String(),
		r.Max().String(),
		r.Quantile(90).String(),
		r.Quantile(95).String(),
		r.Quantile(99).String(),
	)
}

func (r Durs) Avg() time.Duration {
	var avg time.Duration
	for _, dur := range r {
		avg += dur
	}

	return time.Duration(int(avg) / len(r))
}

func (r Durs) Median() time.Duration {
	return r[len(r)/2]
}

func (r Durs) Max() time.Duration {
	return r[len(r)-1]
}

func (r Durs) Min() time.Duration {
	return r[0]
}

func (r Durs) Quantile(p float64) time.Duration {
	idx := int(p * float64(len(r)) / 100)
	return r[idx]
}

func (r Durs) Variance() float64 {
	var sum_ time.Duration
	for _, dur := range r {
		sum_ += dur
	}

	mean := float64(sum_) / float64(len(r))

	var sumDif float64
	for _, dur := range r {
		sumDif += math.Pow(float64(dur)-mean, 2)
	}

	return sumDif / float64(len(r))
}

func (r Durs) SE() float64 {
	return math.Sqrt(r.Variance())
}

func (r Durs) AvgWithSE() [3]time.Duration {
	avg := r.Avg()
	se := r.SE()

	return [3]time.Duration{avg - time.Duration(se), avg, avg + time.Duration(se)}
}
