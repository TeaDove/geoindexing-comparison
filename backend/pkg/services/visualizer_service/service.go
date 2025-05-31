package visualizer_service

import (
	"geoindexing_comparison/pkg/generator"
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/index"
	"geoindexing_comparison/pkg/services/builder_service"
	"sync"

	"github.com/pkg/errors"
)

type Service struct {
	builderService *builder_service.Service

	mu sync.RWMutex

	generatorObj generator.Impl
	indexObj     index.Impl

	points geo.Points
}

func NewService(builderService *builder_service.Service) (*Service, error) {
	r := Service{builderService: builderService}

	err := r.SetVisualizer(&NewVisualizerInput{
		Amount:    10_000,
		Index:     r.builderService.Indexes[0].Info.ShortName,
		Generator: r.builderService.Generators[0].Info.ShortName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to set visualizer")
	}

	return &r, nil
}
