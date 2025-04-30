package main

import (
	"geoindexing_comparison/backend/presentation"
	"geoindexing_comparison/backend/repository"
	"geoindexing_comparison/backend/service"
	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
	"github.com/teadove/teasutils/utils/notify_utils"
	"github.com/teadove/teasutils/utils/settings_utils"
	"os"
	"runtime/pprof"
)

func withProfiler() error {
	file, err := os.Create(settings_utils.ServiceSettings.Prof.ResultFilename)
	if err != nil {
		return errors.Wrap(err, "could not open result file")
	}

	err = pprof.StartCPUProfile(file)
	if err != nil {
		return errors.Wrap(err, "could not start CPU profile")
	}

	notify_utils.RunOnInterruptAndExit(func() {
		pprof.StopCPUProfile()
	})

	return nil
}

func main() {
	ctx := logger_utils.NewLoggedCtx()
	err := withProfiler()
	if err != nil {
		panic(errors.Wrap(err, "could not start profiler"))
	}

	runsRepository, err := repository.NewRepository(ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize repository"))
	}

	app := presentation.NewPresentation(service.NewRunner(ctx, runsRepository))

	err = app.Run(ctx, "0.0.0.0:8000")
	if err != nil {
		panic(errors.Wrap(err, "failed to start server"))
	}
}
