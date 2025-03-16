package main

import (
	"geoindexing_comparison/service/cases"
	"geoindexing_comparison/service/presentation"
	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func main() {
	ctx := logger_utils.NewLoggedCtx()

	app := presentation.NewPresentation(cases.NewRunner(ctx))

	err := app.Run(ctx, "0.0.0.0:80")
	if err != nil {
		panic(errors.Wrap(err, "failed to start server"))
	}
}
