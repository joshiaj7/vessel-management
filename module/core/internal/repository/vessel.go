package repository

//go:generate mockgen -source vessel.go -destination mock/vessel.go

import (
	"context"

	"gorm.io/gorm"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

type VesselRepository interface {
	LockVessel(ctx context.Context, eObj *entity.Vessel) error
	ListVessels(ctx context.Context, params *param.ListVessels) ([]*entity.Vessel, *util.OffsetPagination, error)
	GetVessel(ctx context.Context, params *param.GetVessel) (*entity.Vessel, error)
	CreateVessel(ctx context.Context, params *param.CreateVessel) (*entity.Vessel, error)
	UpdateVessel(ctx context.Context, eObj *entity.Vessel, params *param.UpdateVessel) error
}

type vesselRepository struct {
	databaseName string
	database     *gorm.DB
}

func NewVesselRepository(databaseName string, database *gorm.DB) *vesselRepository {
	return &vesselRepository{
		databaseName: databaseName,
		database:     database,
	}
}

func (r *vesselRepository) LockVessel(ctx context.Context, eObj *entity.Vessel) error {
	return nil
}

func (r *vesselRepository) CreateVessel(ctx context.Context, params *param.CreateVessel) (result *entity.Vessel, err error) {
	return result, nil
}

func (r *vesselRepository) ListVessels(ctx context.Context, params *param.ListVessels) (result []*entity.Vessel, meta *util.OffsetPagination, err error) {
	return result, nil, nil
}

func (r *vesselRepository) GetVessel(ctx context.Context, params *param.GetVessel) (result *entity.Vessel, err error) {
	return result, nil
}

func (r *vesselRepository) UpdateVessel(ctx context.Context, eObj *entity.Vessel, params *param.UpdateVessel) (err error) {
	return nil
}
