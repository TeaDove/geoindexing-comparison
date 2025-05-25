package main

import (
	"geoindexing_comparison/backend/presentations/manager_presentation"
	"geoindexing_comparison/backend/repositories/manager_repository"
	"geoindexing_comparison/backend/services/builder_service"
	"geoindexing_comparison/backend/services/manager_service"
	"geoindexing_comparison/backend/services/visualizer_service"

	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func main() {
	ctx := logger_utils.NewLoggedCtx()

	runsRepository, err := manager_repository.NewRepository(ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize repository"))
	}

	builderService := builder_service.NewService()
	managerService := manager_service.NewService(builderService, runsRepository)

	visualizerService, err := visualizer_service.NewService(builderService)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize visualizer service"))
	}

	app := manager_presentation.NewPresentation(managerService, visualizerService, builderService)

	err = app.Run("0.0.0.0:8000")
	if err != nil {
		panic(errors.Wrap(err, "failed to start server"))
	}
}
