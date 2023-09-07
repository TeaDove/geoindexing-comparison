package run

import (
	"fmt"
	"geoindexing_comparison/cases"
	"geoindexing_comparison/generator"
	"testing"
)

func BenchmarkUnit_Ok(b *testing.B) {
	collections := cases.NewCollections()
	input := &cases.KNNInput{
		Amount: 1000,
		Points: generator.DefaultGenerator.GeneratePointsDefaultAmount(),
		Origin: generator.DefaultGenerator.GeneratePoint(),
	}
	for _, col := range collections {
		runCase := cases.Case{
			Collection: col,
		}

		b.Run(fmt.Sprintf("KNN_%s", col.Name()), func(bb *testing.B) {
			for i := 0; i < bb.N; i++ {
				runCase.KNN(bb, input)
			}
		})
	}
}
