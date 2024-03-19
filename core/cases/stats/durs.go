package stats

import (
	"fmt"
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
