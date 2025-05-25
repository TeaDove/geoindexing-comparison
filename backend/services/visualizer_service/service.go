package visualizer_service

import (
	"geoindexing_comparison/backend/generator"
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"geoindexing_comparison/backend/services/builder_service"
	"github.com/pkg/errors"
	"sync"
)

type Service struct {
	builderService *builder_service.Service

	mu sync.RWMutex

	generator     generator.Impl
	generatorInfo generator.Info
	index         index.Impl
	indexInfo     index.Info

	points geo.Points
}

func NewService(builderService *builder_service.Service) (*Service, error) {
	r := Service{builderService: builderService}

	err := r.SetVisualizer(&NewVisualizerInput{Amount: 10_000, Index: r.builderService.Indexes[0].Info.ShortName})
	if err != nil {
		return nil, errors.Wrap(err, "failed to set visualizer")
	}

	return &r, nil
}
