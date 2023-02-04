package repository

//go:generate mockgen -source voyage.go -destination mock/voyage.go

import (
	"context"

	"gorm.io/gorm"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

type VoyageRepository interface {
	CreateVoyage(ctx context.Context, params *param.CreateVoyage) (*entity.Voyage, error)
	ListVoyages(ctx context.Context, params *param.ListVoyages) ([]*entity.Voyage, *util.OffsetPagination, error)
	GetVoyage(ctx context.Context, params *param.GetVoyage) (*entity.Voyage, error)
	UpdateVoyage(ctx context.Context, eObj *entity.Voyage) (*entity.Voyage, error)
}

type voyageRepository struct {
	databaseName string
	database     *gorm.DB
}

func NewVoyageRepository(databaseName string, database *gorm.DB) *voyageRepository {
	return &voyageRepository{
		databaseName: databaseName,
		database:     database,
	}
}

// nolint gochecknoglobals
var (
	VoyageTable   = "voyages"
	VoyageColumns = []string{
		"id",
		"vessel_id",
		"source",
		"destination",
		"current_location",
		"state",
		"docked_at",
		"departed_at",
		"arrived_at",
		"created_at",
		"updated_at",
	}
)

func (r *voyageRepository) CreateVoyage(ctx context.Context, params *param.CreateVoyage) (result *entity.Voyage, err error) {
	return result, nil
}

func (r *voyageRepository) ListVoyages(ctx context.Context, params *param.ListVoyages) (result []*entity.Voyage, meta *util.OffsetPagination, err error) {
	return result, nil, nil
}

func (r *voyageRepository) GetVoyage(ctx context.Context, params *param.GetVoyage) (result *entity.Voyage, err error) {
	return result, nil
}

func (r *voyageRepository) UpdateVoyage(ctx context.Context, eObj *entity.Voyage) (result *entity.Voyage, err error) {
	return result, nil
}
