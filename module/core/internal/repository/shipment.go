package repository

//go:generate mockgen -source shipment.go -destination mock/shipment.go

import (
	"context"

	"gorm.io/gorm"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

type ShipmentRepository interface {
	CreateShipment(ctx context.Context, params *param.CreateShipment) (*entity.Shipment, error)
	ListShipments(ctx context.Context, params *param.ListShipments) ([]*entity.Shipment, *util.OffsetPagination, error)
	GetShipment(ctx context.Context, params *param.GetShipment) (*entity.Shipment, error)
	UpdateShipment(ctx context.Context, eObj *entity.Shipment) (*entity.Shipment, error)
}

type shipmentRepository struct {
	databaseName string
	database     *gorm.DB
}

func NewShipmentRepository(databaseName string, database *gorm.DB) *shipmentRepository {
	return &shipmentRepository{
		databaseName: databaseName,
		database:     database,
	}
}

// nolint gochecknoglobals
var (
	ShipmentTable   = "shipments"
	ShipmentColumns = []string{
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

func (r *shipmentRepository) CreateShipment(ctx context.Context, params *param.CreateShipment) (result *entity.Shipment, err error) {
	return result, nil
}

func (r *shipmentRepository) ListShipments(ctx context.Context, params *param.ListShipments) (result []*entity.Shipment, meta *util.OffsetPagination, err error) {
	return result, nil, nil
}

func (r *shipmentRepository) GetShipment(ctx context.Context, params *param.GetShipment) (result *entity.Shipment, err error) {
	return result, nil
}

func (r *shipmentRepository) UpdateShipment(ctx context.Context, eObj *entity.Shipment) (result *entity.Shipment, err error) {
	return result, nil
}
