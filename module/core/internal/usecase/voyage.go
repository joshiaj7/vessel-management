package usecase

//go:generate mockgen -source voyage.go -destination mock/voyage.go

import (
	"context"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/internal/repository"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

type VoyageUsecase interface {
	ListVoyages(ctx context.Context, params *param.ListVoyages) ([]*entity.Voyage, *util.OffsetPagination, error)
	GetVoyage(ctx context.Context, params *param.GetVoyage) (*entity.Voyage, error)
	CreateVoyage(ctx context.Context, params *param.CreateVoyage) (*entity.Voyage, error)
	UpdateVoyage(ctx context.Context, params *param.UpdateVoyage) (*entity.Voyage, error)
}

type voyageUsecaseRepository struct {
	voyage repository.VoyageRepository
	vessel repository.VesselRepository
}

type voyageUsecase struct {
	repository *voyageUsecaseRepository
}

func NewVoyageUsecase(
	voyageRepository repository.VoyageRepository,
	vesselRepository repository.VesselRepository,
) *voyageUsecase {
	return &voyageUsecase{
		repository: &voyageUsecaseRepository{
			voyage: voyageRepository,
			vessel: vesselRepository,
		},
	}
}

// TODO: WIP
func (u *voyageUsecase) CreateVoyage(ctx context.Context, params *param.CreateVoyage) (result *entity.Voyage, err error) {
	result, err = u.repository.voyage.CreateVoyage(ctx, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *voyageUsecase) ListVoyages(ctx context.Context, params *param.ListVoyages) (result []*entity.Voyage, pagination *util.OffsetPagination, err error) {
	result, pagination, err = u.repository.voyage.ListVoyages(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	return result, pagination, nil
}

func (u *voyageUsecase) GetVoyage(ctx context.Context, params *param.GetVoyage) (result *entity.Voyage, err error) {
	result, err = u.repository.voyage.GetVoyage(ctx, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *voyageUsecase) UpdateVoyage(ctx context.Context, params *param.UpdateVoyage) (result *entity.Voyage, err error) {
	voyage, err := u.repository.voyage.GetVoyage(ctx, &param.GetVoyage{ID: params.ID})
	if err != nil {
		return nil, err
	}

	result, err = u.repository.voyage.UpdateVoyage(ctx, voyage)
	if err != nil {
		return nil, err
	}

	return result, nil
}
