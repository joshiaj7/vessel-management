package config

import (
	"log"

	"github.com/joshiaj7/vessel-management/internal/util"
	coreConfig "github.com/joshiaj7/vessel-management/module/core/config"
	"github.com/julienschmidt/httprouter"
)

type GatewayConfig struct {
	GatewayHost string `envconfig:"GATEWAY_HOST"`
	ServiceConfig
}

func NewGatewayServer() (router *httprouter.Router, err error) {
	cfg, err := initGatewayConfig()
	if err != nil {
		return nil, util.ErrorWrap(err)
	}

	return registerGatewayCore(cfg)
}

func initGatewayConfig() (cfg GatewayConfig, err error) {
	cfg, err = loadGatewayConfig()
	if err != nil {
		log.Fatalf("Load Gateway Config Failed: %v", err)
		return cfg, util.ErrorWrap(err)
	}

	db, err := NewDB(cfg.DatabaseConfig)
	if err != nil {
		log.Fatalf("Create new DB Failed: %v", err)
		return cfg, util.ErrorWrap(err)
	}
	cfg.Database = db

	return cfg, nil
}

func registerGatewayCore(cfg GatewayConfig) (router *httprouter.Router, err error) {
	coreCfg, err := loadCoreConfig()
	if err != nil {
		log.Fatalf("Load Core Config Failed: %v", err)
		return nil, util.ErrorWrap(err)
	}

	router = coreConfig.RegisterCoreGateway(&coreConfig.GatewayConfig{
		Config:   coreCfg,
		Database: cfg.Database,
	})
	return router, nil
}
