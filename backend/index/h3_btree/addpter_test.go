package h3_btree

import (
	"geoindexing_comparison/backend/generator"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teadove/teasutils/utils/test_utils"
)

func TestKNNOk(t *testing.T) {
	t.Parallel()

	points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, 1000)
	collection := Factory(4)()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	result, _ := collection.KNNTimed(origin, 100)
	assert.Len(t, result, 100)
}

func TestRangeSearchOk(t *testing.T) {
	t.Parallel()

	points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, 1000)
	collection := Factory(5)()
	collection.FromArray(points)

	points, _ = collection.BBoxTimed(points.GetRandomPoint(), points.GetRandomPoint())
	test_utils.Pprint(len(points))
}
