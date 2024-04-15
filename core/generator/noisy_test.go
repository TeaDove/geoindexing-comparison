package generator

import (
	"fmt"
	"testing"
	"time"

	"github.com/KEINOS/go-noise"
	"github.com/stretchr/testify/require"
)

func TestUnit_NoisyGenerator_Point_Ok(t *testing.T) {
	t.Parallel()

	n, err := noise.New(noise.Perlin, time.Now().Unix())
	require.NoError(t, err)

	for i := -1.0; i < 1; i += 0.1 {
		for j := -1.0; j < 1; j += 0.1 {
			v := n.Eval64(i, j)

			fmt.Printf("(%f, %f): %f\n", i, j, v)
		}
	}
}
