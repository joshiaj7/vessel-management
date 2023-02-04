package config

import (
	"github.com/joshiaj7/vessel-management/module/core/internal/usecase"
)

type CoreUsecase struct {
	Voyage usecase.VoyageUsecase
	Vessel usecase.VesselUsecase
}

func RegisterCoreUsecase(repository *CoreRepository) *CoreUsecase {
	voyageUsecase := usecase.NewVoyageUsecase(
		repository.Voyage,
		repository.Vessel,
	)

	vesselUsecase := usecase.NewVesselUsecase(
		repository.Vessel,
	)

	return &CoreUsecase{
		Voyage: voyageUsecase,
		Vessel: vesselUsecase,
	}
}
