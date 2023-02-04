package config

import (
	"github.com/joshiaj7/vessel-management/module/core/internal/repository"
)

type CoreRepository struct {
	Shipment repository.ShipmentRepository
	Vessel   repository.VesselRepository
}

func RegisterCoreRepository(cfg *GatewayConfig) *CoreRepository {
	shipmentRepository := repository.NewShipmentRepository(
		cfg.Config.CoreDatabaseName,
		cfg.Database,
	)

	vesselRepository := repository.NewVesselRepository(
		cfg.Config.CoreDatabaseName,
		cfg.Database,
	)

	return &CoreRepository{
		Shipment: shipmentRepository,
		Vessel:   vesselRepository,
	}
}
