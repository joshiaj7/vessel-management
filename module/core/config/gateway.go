package config

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/julienschmidt/httprouter"

	"github.com/joshiaj7/vessel-management/module/core/internal/handler"
	"github.com/joshiaj7/vessel-management/module/core/internal/repository"
	"github.com/joshiaj7/vessel-management/module/core/internal/usecase"
)

type GatewayConfig struct {
	Config   CoreConfig
	Database *gorm.DB
}

func RegisterCoreHandler(cfg *GatewayConfig) error {

	// Repository

	shipmentRepository := repository.NewShipmentRepository(
		cfg.Config.CoreDatabaseName,
		cfg.Database,
	)

	vesselRepository := repository.NewVesselRepository(
		cfg.Config.CoreDatabaseName,
		cfg.Database,
	)

	// Usecase

	shipmentUsecase := usecase.NewShipmentUsecase(
		shipmentRepository,
		vesselRepository,
	)

	vesselUsecase := usecase.NewVesselUsecase(
		vesselRepository,
	)

	// Handler

	shipmentHandler := handler.NewShipmentHandler(
		shipmentUsecase,
	)

	vesselHandler := handler.NewVesselHandler(
		vesselUsecase,
	)

	router := httprouter.New()
	shipmentHandler.Register(router)
	vesselHandler.Register(router)

	router.HandlerFunc("GET", "/healthz", healthz)

	return nil
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}

// func notFound(w http.ResponseWriter, _ *http.Request) {
// 	meta := response.MetaInfo{HTTPStatus: 404}
// 	res := response.BuildSuccess(nil, "not found", meta)

// 	_ = response.Write(w, res, meta.HTTPStatus)
// }
