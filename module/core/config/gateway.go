package config

import (
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type GatewayConfig struct {
	Config   CoreConfig
	Database *gorm.DB
	Router   *httprouter.Router
}

func RegisterCoreGateway(cfg *GatewayConfig) {

	repositories := RegisterCoreRepository(cfg)
	usecases := RegisterCoreUsecase(repositories)
	RegisterCoreHandler(cfg.Router, usecases)
}
