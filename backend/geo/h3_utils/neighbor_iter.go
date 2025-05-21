package h3_utils

import (
	"github.com/pkg/errors"
	"github.com/uber/h3-go/v4"
	"iter"
)

func GridDiskInf(v h3.Cell) iter.Seq[[]h3.Cell] {
	return func(yield func([]h3.Cell) bool) {
		var k = 0
		for {
			cells, err := h3.GridDiskDistances(v, k)
			if err != nil {
				panic(errors.Wrap(err, "failed to calculate neighbors"))
			}

			if !yield(cells[len(cells)-1]) {
				return
			}
			k++
		}
	}
}
