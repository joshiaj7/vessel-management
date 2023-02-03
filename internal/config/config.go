package config

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // blank import for initializing mysql driver

	sqlxtr "go.bukalapak.io/buka20/core-service/go/sentry/tracing/sqlx"
)

type ServiceConfig struct {
	Environment            string `envconfig:"ENVIRONMENT" default:"dev"`
	AuthDomainHost         string `envconfig:"AUTH_DOMAIN_HOST"`
	AuthDomainEmployeeHost string `envconfig:"AUTH_DOMAIN_EMPLOYEE_HOST"`

	DatabaseConfig DatabaseConfig `envconfig:"DB"`
	// KafkaConfig    KafkaConfig    `envconfig:"KAFKA"`

	Database *sqlxtr.DB `ignored:"true"`
}

type DatabaseConfig struct {
	Driver      string `required:"true" envconfig:"DRIVER"`
	Host        string `required:"true" envconfig:"HOST"`
	Port        int    `required:"true" envconfig:"PORT"`
	Username    string `required:"true" envconfig:"USERNAME"`
	Password    string `required:"true" envconfig:"PASSWORD"`
	Database    string `required:"true" envconfig:"DATABASE"`
	QueryString string `required:"true" envconfig:"QUERYSTRING"`

	MaxOpenConns    int           `required:"true" envconfig:"MAXOPENCONNS"`
	MaxIdleConns    int           `required:"true" envconfig:"MAXIDLECONNS"`
	ConnMaxLifetime time.Duration `required:"true" envconfig:"MAXLIFETIME"`
	ConnMaxIdletime time.Duration `required:"true" envconfig:"MAXIDLETIME"`
}

func (c *DatabaseConfig) RWDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.QueryString,
	)
}

// type CacheConfig struct {
// 	Host     string `envconfig:"HOST"`
// 	Password string `envconfig:"PASSWORD"`
// }

// type KafkaConfig struct {
// 	ClientID string   `envconfig:"CLIENT_ID"`
// 	Brokers  []string `envconfig:"BROKERS"`
// }

// // loadConfig loads the config for service. It can be called from any application
// // instantiation (e.g. consumer or rpc server).
// func loadConfig() (ServiceConfig, error) {
// 	var cfg ServiceConfig

// 	// load from .env if exists
// 	if _, err := os.Stat(".env"); err == nil {
// 		if err := gotenv.Load(); err != nil {
// 			return cfg, liberr.TracerFromError(err)
// 		}
// 	}

// 	// parse cenvironment variable to config struct using "service" namespace
// 	// to prevent conflict with another modules
// 	err := envconfig.Process("service", &cfg)
// 	if err != nil {
// 		return cfg, liberr.TracerFromError(err)
// 	}
// 	return cfg, nil
// }

// func loadGatewayConfig() (GatewayConfig, error) {
// 	var cfg GatewayConfig

// 	// load from .env if exists
// 	if _, err := os.Stat(".env"); err == nil {
// 		if err := gotenv.Load(); err != nil {
// 			return cfg, liberr.TracerFromError(err)
// 		}
// 	}

// 	// parse environment variable to config struct using "service" namespace
// 	// to prevent conflict with another modules
// 	err := envconfig.Process("service", &cfg)
// 	if err != nil {
// 		return cfg, liberr.TracerFromError(err)
// 	}
// 	return cfg, nil
// }
