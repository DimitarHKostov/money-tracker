package database_credentials_formatter_factory

import (
	"fmt"
	"money-tracker/pkg/database/database_config"
	"money-tracker/pkg/database/database_type"
)

const (
	postgres              = "postgres"
	postgresTemplate      = "host=%s port=%s user=%s password=%s sslmode=disable"
	dbTypeNotSupportedMsg = "db type not supported"
)

type DatabaseCredentialsFormatterFactory struct {
}

func (dcff *DatabaseCredentialsFormatterFactory) ProduceFormattedCredentials(databaseConfig database_config.DatabaseConfig) string {
	switch databaseConfig.DatabaseType {
	case database_type.Postgres:
		dbConnectionString := fmt.Sprintf(postgresTemplate, databaseConfig.Host, databaseConfig.Port, databaseConfig.Username, databaseConfig.Password)
		return dbConnectionString
	default:
		return ""
	}
}
