package main

import (
	"geoindexing_comparison/backend/services/builder_service"
	"geoindexing_comparison/backend/services/worker_service"
	"geoindexing_comparison/backend/suppliers/manager_supplier"
)

func main() {
	builderService := builder_service.NewService()
	managerSupplier := manager_supplier.NewSupplier()

	workerService := worker_service.NewService(builderService, managerSupplier)
	workerService.Job()
}
