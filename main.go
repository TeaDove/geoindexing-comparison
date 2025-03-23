package main

import (
	"embed"
	"geoindexing_comparison/service/cases"
	"geoindexing_comparison/service/presentation"
	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
	"github.com/teadove/teasutils/utils/settings_utils"
	"net/http"
)

//go:embed frontend/*
var frontend embed.FS

func getFrontend() http.FileSystem {
	if settings_utils.ServiceSettings.Release {
		return http.FS(frontend)
	}

	return http.Dir("./frontend")
}

func main() {
	ctx := logger_utils.NewLoggedCtx()

	app := presentation.NewPresentation(cases.NewRunner(ctx), getFrontend())

	err := app.Run(ctx, "0.0.0.0:80")
	if err != nil {
		panic(errors.Wrap(err, "failed to start server"))
	}
}
