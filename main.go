package main

import (
	"geoindexing_comparison/service/cases"
	"geoindexing_comparison/service/presentation"
	"geoindexing_comparison/service/repository"
	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func main() {
	ctx := logger_utils.NewLoggedCtx()

	runsRepository, err := repository.NewRepository(ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize repository"))
	}

	app := presentation.NewPresentation(cases.NewRunner(ctx), runsRepository)

	err = app.Run(ctx, "0.0.0.0:8000")
	if err != nil {
		panic(errors.Wrap(err, "failed to start server"))
	}
}
