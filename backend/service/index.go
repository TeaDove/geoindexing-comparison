package service

import (
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"github.com/pkg/errors"
	"time"
)

type Visualizer struct {
	Generator     generator.Impl `json:"-"`
	GeneratorInfo generator.Info `json:"generatorInfo"`
	Index         index.Impl     `json:"-"`
	IndexInfo     index.Info     `json:"indexInfo"`
}

type NewVisualizerInput struct {
	Amount    uint64 `json:"amount"`
	IndexName string `json:"indexName"`
}

func (r *Service) SetVisualizer(input *NewVisualizerInput) (*Visualizer, error) {
	gen := generator.AllGenerators()[0]
	idx, ok := r.NameToIndex[input.IndexName]
	if !ok {
		return nil, errors.Errorf("index not found: %s", input.IndexName)
	}

	visualizer := Visualizer{GeneratorInfo: gen.Info, Generator: gen.Builder(), Index: idx.Builder(), IndexInfo: idx.Info}

	points := visualizer.Generator.Points(&generator.DefaultInput, input.Amount)

	visualizer.Index.FromArray(points)
	r.Visualizer = visualizer
	return &visualizer, nil
}

func (r *Visualizer) GetPoints() geo.Points {
	return r.Index.ToArray()
}

type KNNInput struct {
	Point geo.Point `json:"point"`
	N     uint64    `json:"n"`
}

func (r *Visualizer) KNN(input *KNNInput) (geo.Points, time.Duration) {
	return r.Index.KNNTimed(input.Point, input.N)
}

type RangeSearchInput struct {
	Point  geo.Point `json:"point"`
	Radius float64   `json:"radius"`
}

func (r *Visualizer) RangeSearch(input *RangeSearchInput) (geo.Points, time.Duration) {
	return r.Index.RangeSearchTimed(input.Point, input.Radius)
}
