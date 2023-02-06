package config

import (
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type GatewayConfig struct {
	Config   CoreConfig
	Database *gorm.DB
}

func RegisterCoreGateway(cfg *GatewayConfig) *httprouter.Router {

	repositories := RegisterCoreRepository(cfg)
	usecases := RegisterCoreUsecase(repositories)
	return RegisterCoreHandler(usecases)
}
