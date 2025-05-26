package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_QualifiedAvg(t *testing.T) {
	arr := Array[int]{0, 2, 3, 4, 11}

	assert.InEpsilon(t, 4.0, arr.Avg(), 0.00001)
	assert.InEpsilon(t, 3.0, arr.QualifiedAvg(), 0.00001)
	assert.Equal(t, 3, arr.Median())
}
