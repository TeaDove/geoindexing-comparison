package manager_service

import (
	"geoindexing_comparison/pkg/repositories/manager_repository"
	"geoindexing_comparison/pkg/schemas"
	"geoindexing_comparison/pkg/services/builder_service"
	"sync"
)

type Service struct {
	builderService *builder_service.Service

	repository *manager_repository.Repository

	currentRun  *manager_repository.Run
	jobMu       sync.Mutex
	jobIdx      int
	jobs        []schemas.Job
	allJobsDone chan struct{}
}

func NewService(builderService *builder_service.Service, repository *manager_repository.Repository) *Service {
	r := Service{builderService: builderService, repository: repository}

	go r.initRunner()

	return &r
}
