package usecase

import (
	"context"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/internal/repository"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

type ShipmentUsecase interface {
	ListShipments(ctx context.Context, params *param.ListShipments) ([]*entity.Shipment, *util.OffsetPagination, error)
	GetShipment(ctx context.Context, params *param.GetShipment) (*entity.Shipment, error)
	CreateShipment(ctx context.Context, params *param.CreateShipment) (*entity.Shipment, error)
	UpdateShipment(ctx context.Context, eObj *entity.Shipment) (*entity.Shipment, error)
}

type shipmentUsecaseRepository struct {
	shipment repository.ShipmentRepository
	vessel   repository.VesselRepository
}

type shipmentUsecase struct {
	repository *shipmentUsecaseRepository
}

func NewShipmentUsecase(
	shipmentRepository repository.ShipmentRepository,
	vesselRepository repository.VesselRepository,
) *shipmentUsecase {
	return &shipmentUsecase{
		repository: &shipmentUsecaseRepository{
			shipment: shipmentRepository,
			vessel:   vesselRepository,
		},
	}
}

func (u *shipmentUsecase) CreateShipment(ctx context.Context, params *param.CreateShipment) (result *entity.Shipment, err error) {
	result, err = u.repository.shipment.CreateShipment(ctx, params)
	if err != nil {
		return nil, util.ErrorWrap(err)
	}

	return result, nil
}

func (u *shipmentUsecase) ListShipments(ctx context.Context, params *param.ListShipments) (result []*entity.Shipment, pagination *util.OffsetPagination, err error) {
	result, pagination, err = u.repository.shipment.ListShipments(ctx, params)
	if err != nil {
		return nil, nil, util.ErrorWrap(err)
	}

	return result, pagination, nil
}

func (u *shipmentUsecase) GetShipment(ctx context.Context, params *param.GetShipment) (result *entity.Shipment, err error) {
	result, err = u.repository.shipment.GetShipment(ctx, params)
	if err != nil {
		return nil, util.ErrorWrap(err)
	}

	return result, nil
}

func (u *shipmentUsecase) UpdateShipment(ctx context.Context, eObj *entity.Shipment) (result *entity.Shipment, err error) {
	result, err = u.repository.shipment.UpdateShipment(ctx, eObj)
	if err != nil {
		return nil, util.ErrorWrap(err)
	}

	return result, nil
}
