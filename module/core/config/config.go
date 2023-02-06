package config

type CoreConfig struct {
	CoreDatabaseName string `required:"true" envconfig:"DATABASE_NAME"`
}
