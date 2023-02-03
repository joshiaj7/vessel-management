package config

import (
	"context"
	"net/http"

	"github.com/joshiaj7/vessel-management/internal/util"
)

type GatewayConfig struct {
	GatewayHost              string   `envconfig:"GATEWAY_HOST"`
	AllowedCredentialOrigins []string `envconfig:"ALLOWED_CREDENTIAL_ORIGINS"`
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

func NewGatewayServer() (*GatewayServer, error) {
	return nil, nil
}
