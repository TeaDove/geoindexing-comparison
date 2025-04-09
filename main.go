package main

import (
	"geoindexing_comparison/backend/presentation"
	"geoindexing_comparison/backend/repository"
	"geoindexing_comparison/backend/service"
	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func main() {
	ctx := logger_utils.NewLoggedCtx()

	runsRepository, err := repository.NewRepository(ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize repository"))
	}

	app := presentation.NewPresentation(service.NewRunner(ctx, runsRepository), runsRepository)

	err = app.Run(ctx, "0.0.0.0:8000")
	if err != nil {
		panic(errors.Wrap(err, "failed to start server"))
	}
}
