package worker_service

import (
	"geoindexing_comparison/pkg/services/builder_service"
	"geoindexing_comparison/pkg/suppliers/manager_supplier"
)

type Service struct {
	builderService *builder_service.Service

	managerSupplier *manager_supplier.Supplier
}

func NewService(builderService *builder_service.Service, managerSupplier *manager_supplier.Supplier) *Service {
	return &Service{builderService: builderService, managerSupplier: managerSupplier}
}
