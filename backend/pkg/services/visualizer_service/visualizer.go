package visualizer_service

import (
	"geoindexing_comparison/pkg/generator"
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/helpers"
	"time"

	"github.com/pkg/errors"
)

type NewVisualizerInput struct {
	Generator string `json:"generator"`
	Amount    int    `json:"amount"`
	Index     string `json:"index"`
}

func (r *Service) SetVisualizer(input *NewVisualizerInput) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if input.Index != "" {
		idx, ok := r.builderService.IndexMap[input.Index]
		if !ok {
			return errors.Errorf("index not found: %s", input.Index)
		}

		r.indexObj = idx.Builder()
		r.indexObj.FromArray(r.points)
	}

	if input.Amount != 0 || input.Generator != "" {
		if input.Generator != "" {
			gen, ok := r.builderService.GeneratorMap[input.Generator]
			if !ok {
				return errors.Errorf("generator not found: %s", input.Generator)
			}

			r.generatorObj = gen.Builder(helpers.RNG())
		}

		r.points = r.generatorObj.Points(&generator.DefaultInput, input.Amount)
		r.indexObj.FromArray(r.points)
	}

	return nil
}

func (r *Service) GetPoints() geo.Points {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.points
}

type KNNInput struct {
	Point geo.Point `json:"point"`
	N     int       `json:"n"`
}

func (r *Service) KNN(input *KNNInput) (geo.Points, time.Duration) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.indexObj.KNNTimed(input.Point, input.N)
}

type BBoxInput struct {
	BottomLeft geo.Point `json:"bottomLeft"`
	UpperRight geo.Point `json:"upperRight"`
}

func (r *Service) BBox(input *BBoxInput) (geo.Points, time.Duration) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.indexObj.BBoxTimed(input.BottomLeft, input.UpperRight)
}
