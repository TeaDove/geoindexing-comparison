package main

import (
	"geoindexing_comparison/backend/presentations/manager_presentation"
	"geoindexing_comparison/backend/repositories/manager_repository"
	manager_service "geoindexing_comparison/backend/services/worker_service"
	"os"
	"runtime/pprof"

	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/logger_utils"
	"github.com/teadove/teasutils/utils/notify_utils"
)

func withProfiler() error {
	file, err := os.Create("cpu.prof")
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

	runsRepository, err := manager_repository.NewRepository(ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize repository"))
	}

	runner := manager_service.NewService(ctx, runsRepository)

	app := manager_presentation.NewPresentation(runner)

	err = app.Run(ctx, "0.0.0.0:8000")
	if err != nil {
		panic(errors.Wrap(err, "failed to start server"))
	}
}
