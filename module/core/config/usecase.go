package config

import (
	"github.com/joshiaj7/vessel-management/module/core/internal/usecase"
)

type CoreUsecase struct {
	Shipment usecase.ShipmentUsecase
	Vessel   usecase.VesselUsecase
}

func RegisterCoreUsecase(repository *CoreRepository) *CoreUsecase {
	shipmentUsecase := usecase.NewShipmentUsecase(
		repository.Shipment,
		repository.Vessel,
	)

	vesselUsecase := usecase.NewVesselUsecase(
		repository.Vessel,
	)

	return &CoreUsecase{
		Shipment: shipmentUsecase,
		Vessel:   vesselUsecase,
	}
}
