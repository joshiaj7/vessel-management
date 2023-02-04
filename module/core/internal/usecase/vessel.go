package usecase

import (
	"context"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/internal/repository"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

type VesselUsecase interface {
	CreateVessel(ctx context.Context, params *param.CreateVessel) (*entity.Vessel, error)
	ListVessels(ctx context.Context, params *param.ListVessels) ([]*entity.Vessel, *util.OffsetPagination, error)
	GetVessel(ctx context.Context, params *param.GetVessel) (*entity.Vessel, error)
	UpdateVessel(ctx context.Context, eObj *entity.Vessel) (*entity.Vessel, error)
}

type vesselUsecaseRepository struct {
	vessel repository.VesselRepository
}

type vesselUsecase struct {
	repository vesselUsecaseRepository
}

func NewVesselUsecase(
	vesselRepository repository.VesselRepository,
) *vesselUsecase {
	return &vesselUsecase{
		repository: vesselUsecaseRepository{
			vessel: vesselRepository,
		},
	}
}

func (u *vesselUsecase) CreateVessel(ctx context.Context, params *param.CreateVessel) (result *entity.Vessel, err error) {
	result, err = u.repository.vessel.CreateVessel(ctx, params)
	if err != nil {
		return nil, util.ErrorWrap(err)
	}

	return result, nil
}

func (u *vesselUsecase) ListVessels(ctx context.Context, params *param.ListVessels) (result []*entity.Vessel, pagination *util.OffsetPagination, err error) {
	result, pagination, err = u.repository.vessel.ListVessels(ctx, params)
	if err != nil {
		return nil, nil, util.ErrorWrap(err)
	}

	return result, pagination, nil
}

func (u *vesselUsecase) GetVessel(ctx context.Context, params *param.GetVessel) (result *entity.Vessel, err error) {
	result, err = u.repository.vessel.GetVessel(ctx, params)
	if err != nil {
		return nil, util.ErrorWrap(err)
	}

	return result, nil
}

func (u *vesselUsecase) UpdateVessel(ctx context.Context, eObj *entity.Vessel) (result *entity.Vessel, err error) {
	result, err = u.repository.vessel.UpdateVessel(ctx, eObj)
	if err != nil {
		return nil, util.ErrorWrap(err)
	}

	return result, nil
}
