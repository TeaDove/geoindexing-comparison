package main

import (
	"geoindexing_comparison/pkg/helpers"
	"geoindexing_comparison/pkg/presentations/manager_presentation"
	"geoindexing_comparison/pkg/repositories/manager_repository"
	"geoindexing_comparison/pkg/services/builder_service"
	"geoindexing_comparison/pkg/services/manager_service"
	"geoindexing_comparison/pkg/services/visualizer_service"

	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func main() {
	ctx := logger_utils.NewLoggedCtx()

	runsRepository, err := manager_repository.NewRepository(ctx, helpers.Settings.SqlitePath)
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
