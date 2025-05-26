package main

import (
	"geoindexing_comparison/pkg/services/builder_service"
	"geoindexing_comparison/pkg/services/worker_service"
	"geoindexing_comparison/pkg/suppliers/manager_supplier"
	"os"
	"runtime/pprof"

	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/notify_utils"
)

func withProfiler() {
	file, err := os.Create("cpu.prof")
	if err != nil {
		panic(errors.Wrap(err, "could not open result file"))
	}

	err = pprof.StartCPUProfile(file)
	if err != nil {
		panic(errors.Wrap(err, "could not start CPU profile"))
	}

	notify_utils.RunOnInterruptAndExit(func() {
		pprof.StopCPUProfile()
	})
}

func main() {
	withProfiler()

	builderService := builder_service.NewService()
	managerSupplier := manager_supplier.NewSupplier()

	workerService := worker_service.NewService(builderService, managerSupplier)
	workerService.Job()
}
