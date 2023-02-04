package config

import (
	"github.com/joshiaj7/vessel-management/module/core/internal/repository"
)

type CoreRepository struct {
	Voyage repository.VoyageRepository
	Vessel repository.VesselRepository
}

func RegisterCoreRepository(cfg *GatewayConfig) *CoreRepository {
	voyageRepository := repository.NewVoyageRepository(
		cfg.Config.CoreDatabaseName,
		cfg.Database,
	)

	vesselRepository := repository.NewVesselRepository(
		cfg.Config.CoreDatabaseName,
		cfg.Database,
	)

	return &CoreRepository{
		Voyage: voyageRepository,
		Vessel: vesselRepository,
	}
}
