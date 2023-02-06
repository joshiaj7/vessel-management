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

func NewGatewayServer() (cfg GatewayConfig, err error) {
	cfg, err = initGatewayConfig()
	if err != nil {
		return cfg, util.ErrorWrap(err)
	}

	registerGatewayCore(cfg)
	return cfg, util.ErrorWrap(err)
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

	router := httprouter.New()
	cfg.Router = router

	return cfg, nil
}

func registerGatewayCore(cfg GatewayConfig) (err error) {
	coreCfg, err := loadCoreConfig()
	if err != nil {
		log.Fatalf("Load Core Config Failed: %v", err)
		return util.ErrorWrap(err)
	}

	coreConfig.RegisterCoreGateway(&coreConfig.GatewayConfig{
		Config:   coreCfg,
		Database: cfg.Database,
		Router:   cfg.Router,
	})
	return nil
}
