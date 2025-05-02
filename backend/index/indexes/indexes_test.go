package indexes

import (
	"fmt"
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type TestInput struct {
	InputName string
	Name      string
	Amount    uint64
	Point     geo.Point
	Points    geo.Points
	Index     index.Index
	Generator generator.Generator
}

func makeTestInputs() TestInputs {
	var inputs []TestInput
	for idx := range 4 {
		amount := 100 * uint64(math.Pow(10, float64(idx)))
		for _, gen := range generator.AllGenerators() {
			genObj := gen.Builder()
			points := genObj.Points(&generator.DefaultInput, amount)
			point := genObj.Point(&generator.DefaultInput)

			for _, indexObj := range AllIndexes() {
				inputs = append(inputs, TestInput{
					InputName: fmt.Sprintf("gen:%s-points:%d", gen.Info.ShortName, amount),
					Name:      fmt.Sprintf("gen:%s-points:%d-index:%s", gen.Info.ShortName, amount, indexObj.Info.ShortName),
					Amount:    amount,
					Points:    points,
					Point:     point,
					Index:     indexObj,
					Generator: gen,
				})
			}
		}
	}

	return inputs
}

type TestInputs []TestInput

func (r TestInputs) PerIndex() map[string][]TestInput {
	perIndex := make(map[string][]TestInput)
	for _, input := range r {
		perIndex[input.InputName] = append(perIndex[input.InputName], input)
	}

	return perIndex
}

var testInputs = makeTestInputs()

func TestFromArrayOk(t *testing.T) {
	t.Parallel()

	for _, tt := range testInputs {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			indexObj := tt.Index.Builder()

			indexObj.FromArray(tt.Points)
			actPoints := indexObj.ToArray()

			assert.True(t, actPoints.Equal(tt.Points))
		})
	}
}

func TestStringOk(t *testing.T) {
	t.Parallel()

	for _, tt := range testInputs {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			indexObj := tt.Index.Builder()

			indexObj.FromArray(tt.Points)
			indexObj.String()
		})
	}
}

func TestKNNOk(t *testing.T) {
	t.Parallel()

	for _, inputs := range testInputs.PerIndex() {
		var results []geo.Points
		for _, tt := range inputs {
			t.Run(tt.Name, func(t *testing.T) {
				indexObj := tt.Index.Builder()
				indexObj.FromArray(tt.Points)

				points, _ := indexObj.KNNTimed(tt.Point, 5)

				assert.Len(t, points, 5)
				results = append(results, points)
			})
		}

		results[0].EqualMany(results)
	}
}

func TestRangeSearchOk(t *testing.T) {
	t.Parallel()

	for _, inputs := range testInputs.PerIndex() {
		var results []geo.Points
		for _, tt := range inputs {
			t.Run(tt.Name, func(t *testing.T) {
				indexObj := tt.Index.Builder()
				indexObj.FromArray(tt.Points)

				neiborhs, _ := indexObj.KNNTimed(tt.Point, 5)

				points, _ := indexObj.RangeSearchTimed(tt.Point, tt.Point.DistanceTo(neiborhs[len(neiborhs)-1]))
				assert.GreaterOrEqual(t, len(points), 5)

				results = append(results, points)
			})
		}

		results[0].EqualMany(results)
	}
}
