package cases

import (
	"fmt"
	"sort"
	"time"
)

type Result struct {
	dur []time.Duration
}

func NewResult(dur []time.Duration) Result {
	sort.Slice(dur, func(i, j int) bool {
		return dur[i] < dur[j]
	})

	return Result{dur}
}

func (r *Result) String() string {
	return fmt.Sprintf("avg: %dms, median: %dms, min: %dms, max: %dms, p90: %dms, p95: %dms, p99: %dms",
		r.Avg().Milliseconds(),
		r.Median().Milliseconds(),
		r.Min().Milliseconds(),
		r.Max().Milliseconds(),
		r.Quantile(90).Milliseconds(),
		r.Quantile(95).Milliseconds(),
		r.Quantile(99).Milliseconds(),
	)
}

func (r *Result) Avg() time.Duration {
	var avg time.Duration
	for _, dur := range r.dur {
		avg += dur
	}

	return time.Duration(int(avg) / len(r.dur))
}

func (r *Result) Median() time.Duration {
	return r.dur[len(r.dur)/2]
}

func (r *Result) Max() time.Duration {
	return r.dur[len(r.dur)-1]
}

func (r *Result) Min() time.Duration {
	return r.dur[0]
}

func (r *Result) Quantile(p float64) time.Duration {
	idx := int(p * float64(len(r.dur)) / 100)
	return r.dur[idx]
}
