package visualizer_service

import (
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/helpers"
	"github.com/pkg/errors"
	"time"
)

type NewVisualizerInput struct {
	Amount int    `json:"amount"`
	Index  string `json:"index"`
}

func (r *Service) SetVisualizer(input *NewVisualizerInput) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if input.Index != "" {
		idx, ok := r.builderService.NameToIndex[input.Index]
		if !ok {
			return errors.Errorf("index not found: %s", input.Index)
		}

		r.indexInfo = idx.Info
		r.index = idx.Builder()
		r.index.FromArray(r.points)
	}

	if input.Amount != 0 {
		gen := generator.AllGenerators()[0]

		r.generatorInfo = gen.Info
		r.generator = gen.Builder(helpers.RNG())
		r.points = r.generator.Points(&generator.DefaultInput, input.Amount)
		r.index.FromArray(r.points)
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

	return r.index.KNNTimed(input.Point, input.N)
}

type BBoxInput struct {
	BottomLeft geo.Point `json:"bottomLeft"`
	UpperRight geo.Point `json:"upperRight"`
}

func (r *Service) BBox(input *BBoxInput) (geo.Points, time.Duration) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.index.BBoxTimed(input.BottomLeft, input.UpperRight)
}
