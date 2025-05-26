package indexes

import (
	"fmt"
	"geoindexing_comparison/pkg/generator"
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/helpers"
	"geoindexing_comparison/pkg/index"
	"math"
	"testing"

	"github.com/teadove/teasutils/utils/test_utils"

	"github.com/stretchr/testify/assert"
)

type TestInput struct {
	InputName string
	Name      string
	Amount    int
	Point     geo.Point
	Points    geo.Points
	Index     index.Index
	Generator generator.Generator
}

func makeTestInputs() TestInputs {
	var inputs []TestInput

	for idx := range 4 {
		amount := 100 * int(math.Pow(10, float64(idx)))

		for _, gen := range generator.AllGenerators() {
			genObj := gen.Builder(helpers.RNG())
			points := genObj.Points(&generator.DefaultInput, amount)
			point := genObj.Point(&generator.DefaultInput)

			for _, indexObj := range AllIndexes() {
				inputs = append(inputs, TestInput{
					InputName: fmt.Sprintf("gen:%s-points:%d", gen.Info.ShortName, amount),
					Name: fmt.Sprintf(
						"gen:%s-points:%d-index:%s",
						gen.Info.ShortName,
						amount,
						indexObj.Info.ShortName,
					),
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

var testInputs = makeTestInputs() //nolint: gochecknoglobals // Allowed for tests

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

	for _, inputs := range testInputs.PerIndex() { //nolint: paralleltest // Fails otherwise
		var results []geo.Points

		for _, tt := range inputs {
			t.Run(tt.Name, func(t *testing.T) {
				indexObj := tt.Index.Builder()
				indexObj.FromArray(tt.Points)

				points, _ := indexObj.KNNTimed(tt.Point, 10)

				assert.Len(t, points, 10)
				results = append(results, points)
			})
			test_utils.Pprint(tt.Index.Info.ShortName)
		}

		results[0].EqualMany(results)
	}
}

func TestBBoxOk(t *testing.T) {
	t.Parallel()

	for _, inputs := range testInputs.PerIndex() { //nolint: paralleltest // Fails otherwise
		var results []geo.Points

		for _, tt := range inputs {
			t.Run(tt.Name, func(_ *testing.T) {
				indexObj := tt.Index.Builder()
				indexObj.FromArray(tt.Points)

				knns, _ := indexObj.KNNTimed(tt.Point, 5)
				bottomLeft, upperRight := knns.FindCorners()

				points, _ := indexObj.BBoxTimed(bottomLeft, upperRight)
				results = append(results, points)
			})
		}

		results[0].EqualMany(results)
	}
}
