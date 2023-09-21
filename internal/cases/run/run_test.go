package run

import (
	"testing"
)

func BenchmarkUnit_Ok(b *testing.B) {
	b.Run("a", func(b *testing.B) {
		println("a")
	})
	//collections := cases.NewCollections()
	//inputs := make([]cases.KNNInput, b.N)
	//for i := 0; i <= b.N; i++ {
	//	inputs[i] = cases.KNNInput{
	//		Amount: 1000,
	//		Points: generator.DefaultGenerator.GeneratePointsDefaultAmount(),
	//		Origin: generator.DefaultGenerator.GeneratePoint(),
	//	}
	//}

	//for _, col := range collections {
	//	runCase := cases.Case{
	//		Collection: col,
	//	}
	//
	//	//b.Run(fmt.Sprintf("KNN_%s", col.Name()), func(bb *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		runCase.KNN(b, &inputs[i])
	//	}
	//	//})
	//}
}
