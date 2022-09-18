package database_config

import (
	"money-tracker/pkg/database/database_type"
)

type DatabaseConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseType database_type.DatabaseType
}
