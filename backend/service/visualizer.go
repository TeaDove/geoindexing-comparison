package service

import (
	"context"
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"sync"
	"time"
)

type Visualizer struct {
	mu sync.RWMutex

	generator     generator.Impl
	generatorInfo generator.Info
	index         index.Impl
	indexInfo     index.Info

	points geo.Points
}

type NewVisualizerInput struct {
	Amount uint64 `json:"amount"`
	Index  string `json:"index"`
}

func (r *Service) SetVisualizer(ctx context.Context, input *NewVisualizerInput) (*Visualizer, error) {
	r.Visualizer.mu.Lock()
	defer r.Visualizer.mu.Unlock()

	if input.Index != "" {
		idx, ok := r.NameToIndex[input.Index]
		if !ok {
			return nil, errors.Errorf("index not found: %s", input.Index)
		}

		r.Visualizer.indexInfo = idx.Info
		r.Visualizer.index = idx.Builder()
		r.Visualizer.index.FromArray(r.Visualizer.points)
	}

	if input.Amount != 0 {
		gen := generator.AllGenerators()[0]

		r.Visualizer.generatorInfo = gen.Info
		r.Visualizer.generator = gen.Builder()
		r.Visualizer.points = r.Visualizer.generator.Points(&generator.DefaultInput, input.Amount)
		r.Visualizer.index.FromArray(r.Visualizer.points)
	}

	zerolog.Ctx(ctx).
		Info().
		Int("v", len(r.Visualizer.GetPoints())).
		Msg("visualizer.set")

	return &r.Visualizer, nil
}

func (r *Visualizer) GetPoints() geo.Points {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.points
}

type KNNInput struct {
	Point geo.Point `json:"point"`
	N     uint64    `json:"n"`
}

func (r *Visualizer) KNN(input *KNNInput) (geo.Points, time.Duration) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.index.KNNTimed(input.Point, input.N)
}

type RangeSearchInput struct {
	Point  geo.Point `json:"point"`
	Radius float64   `json:"radius"`
}

func (r *Visualizer) RangeSearch(input *RangeSearchInput) (geo.Points, time.Duration) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.index.RangeSearchTimed(input.Point, input.Radius)
}
