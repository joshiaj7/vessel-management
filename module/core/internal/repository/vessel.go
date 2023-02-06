package repository

//go:generate mockgen -source vessel.go -destination mock/vessel.go

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

var (
	VesselColumnsInsert = []string{
		"owner_id",
		"name",
		"naccs_code",
		"created_at",
		"updated_at",
	}
	VesselColumns = append([]string{"id"}, VesselColumnsInsert...)
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

func (r *vesselRepository) LockVessel(ctx context.Context, eObj *entity.Vessel) (err error) {
	err = r.database.Select(VesselColumns).
		Where("id = ?", eObj.ID).
		Clauses(clause.Locking{Strength: "UPDATE"}).Find(&eObj).Error

	return err
}

func (r *vesselRepository) CreateVessel(ctx context.Context, params *param.CreateVessel) (result *entity.Vessel, err error) {
	timeNow := time.Now()
	vessel := &entity.Vessel{
		OwnerID:   params.OwnerID,
		Name:      params.Name,
		NACCSCode: params.NACCSCode,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	err = r.database.Select(VesselColumnsInsert).Create(vessel).Error
	if err != nil {
		return nil, err
	}

	return vessel, nil
}

func (r *vesselRepository) ListVessels(ctx context.Context, params *param.ListVessels) (result []*entity.Vessel, meta *util.OffsetPagination, err error) {
	vessels := []*entity.Vessel{}

	query := r.database.Select(VesselColumns)
	cquery := r.database.Model(&entity.Vessel{})

	if params.Name != "" {
		query.Where("name LIKE ?", "%"+params.Name+"%")
		cquery.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.OwnerID != 0 {
		query.Where("owner_id = ?", params.OwnerID)
		cquery.Where("owner_id = ?", params.OwnerID)
	}

	err = query.Limit(params.Limit).Offset(params.Offset).Find(&vessels).Error
	if err != nil {
		return nil, nil, err
	}

	var count int64
	err = cquery.Count(&count).Error
	if err != nil {
		return nil, nil, err
	}

	return vessels, util.NewOffsetPagination(params.Limit, params.Offset, int(count)), nil
}

func (r *vesselRepository) GetVessel(ctx context.Context, params *param.GetVessel) (result *entity.Vessel, err error) {
	vessel := &entity.Vessel{ID: params.ID}

	query := r.database.Select(VesselColumns)
	err = query.Find(&vessel).Error
	if err != nil {
		return nil, err
	}

	return vessel, nil
}

func (r *vesselRepository) UpdateVessel(ctx context.Context, eObj *entity.Vessel, params *param.UpdateVessel) (err error) {
	query := r.database.Model(&eObj)
	updatedColumns := map[string]interface{}{}

	if eObj.OwnerID != params.OwnerID && params.OwnerID > 0 {
		eObj.OwnerID = params.OwnerID
		updatedColumns["owner_id"] = params.OwnerID
	}

	if eObj.Name != params.Name && params.Name != "" {
		eObj.Name = params.Name
		updatedColumns["name"] = params.Name
	}

	if eObj.NACCSCode != params.NACCSCode && params.NACCSCode != "" {
		eObj.NACCSCode = params.NACCSCode
		updatedColumns["naccs_code"] = params.NACCSCode
	}

	timeNow := time.Now()
	updatedColumns["updated_at"] = timeNow
	eObj.UpdatedAt = timeNow

	err = query.Updates(updatedColumns).Error

	return err
}
