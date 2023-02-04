package config

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/joshiaj7/vessel-management/internal/util"
	coreConfig "github.com/joshiaj7/vessel-management/module/core/config"
)

type GatewayConfig struct {
	GatewayHost string `envconfig:"GATEWAY_HOST"`
	ServiceConfig
}

type GatewayServer struct {
	Server *http.Server
	Host   string
}

func (gs *GatewayServer) Shutdown(ctx context.Context) error {
	return util.ErrorWrap(gs.Server.Shutdown(ctx))
}

func (gs *GatewayServer) ListenAndServe() error {
	return util.ErrorWrap(gs.Server.ListenAndServe())
}

func NewGatewayServer() (gateway *GatewayServer, err error) {
	cfg, err := initGatewayConfig()
	if err != nil {
		return nil, util.ErrorWrap(err)
	}

	registerGatewayCore(cfg)

	server := &http.Server{
		Addr:              cfg.GatewayHost,
		Handler:           nil,
		ReadHeaderTimeout: 30 * time.Second,
	}

	return &GatewayServer{
		Server: server,
		Host:   cfg.GatewayHost,
	}, nil
}

func initGatewayConfig() (cfg GatewayConfig, err error) {
	cfg, err = loadGatewayConfig()
	if err != nil {
		log.Fatalf("Load Gateway Config Failed: %v", err)
	}

	db, err := NewDB(cfg.DatabaseConfig)
	if err != nil {
		log.Fatalf("Create new DB Failed: %v", err)
	}
	cfg.Database = db

	return cfg, nil
}

func registerGatewayCore(cfg GatewayConfig) {
	coreCfg, err := loadCoreConfig()
	if err != nil {
		log.Fatalf("Register Gateway Core Failed: %v", err)
	}

	coreConfig.RegisterCoreGateway(&coreConfig.GatewayConfig{
		Config:   coreCfg,
		Database: cfg.Database,
	})
}
