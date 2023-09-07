package cases

import (
	"testing"
)

func (r *Case) KNN(b *testing.B, input *KNNInput) {
	// Arrange
	r.Collection.FromArray(input.Points)

	// Task
	b.StartTimer()
	r.Collection.KNN(input.Origin, input.Amount)
	b.StopTimer()
}
