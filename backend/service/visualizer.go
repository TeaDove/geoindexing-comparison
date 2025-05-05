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

	Generator     generator.Impl `json:"-"`
	GeneratorInfo generator.Info `json:"generatorInfo"`
	Index         index.Impl     `json:"-"`
	IndexInfo     index.Info     `json:"indexInfo"`

	Points geo.Points `json:"-"`
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

		r.Visualizer.IndexInfo = idx.Info
		r.Visualizer.Index = idx.Builder()
		r.Visualizer.Index.FromArray(r.Visualizer.Points)
	}

	if input.Amount != 0 {
		gen := generator.AllGenerators()[0]

		r.Visualizer.GeneratorInfo = gen.Info
		r.Visualizer.Generator = gen.Builder()
		r.Visualizer.Points = r.Visualizer.Generator.Points(&generator.DefaultInput, input.Amount)
		r.Visualizer.Index.FromArray(r.Visualizer.Points)
	}

	zerolog.Ctx(ctx).
		Info().
		Interface("v", r.Visualizer).
		Msg("visualizer.set")

	return &r.Visualizer, nil
}

func (r *Visualizer) GetPoints() geo.Points {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.Points
}

type KNNInput struct {
	Point geo.Point `json:"point"`
	N     uint64    `json:"n"`
}

func (r *Visualizer) KNN(input *KNNInput) (geo.Points, time.Duration) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.Index.KNNTimed(input.Point, input.N)
}

type RangeSearchInput struct {
	Point  geo.Point `json:"point"`
	Radius float64   `json:"radius"`
}

func (r *Visualizer) RangeSearch(input *RangeSearchInput) (geo.Points, time.Duration) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.Index.RangeSearchTimed(input.Point, input.Radius)
}
