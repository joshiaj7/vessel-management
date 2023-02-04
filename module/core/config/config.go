package config

import "gorm.io/gorm"

type CoreConfig struct {
	CoreDatabaseName string `required:"true" envconfig:"DATABASE_NAME"`
	Database         *gorm.DB
}
