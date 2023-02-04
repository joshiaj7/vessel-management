package config

import (
	"gorm.io/gorm"
)

type GatewayConfig struct {
	Config   CoreConfig
	Database *gorm.DB
}

func RegisterCoreGateway(cfg *GatewayConfig) error {

	repositories := RegisterCoreRepository(cfg)
	usecases := RegisterCoreUsecase(repositories)
	RegisterCoreHandler(usecases)

	return nil
}
