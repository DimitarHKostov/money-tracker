package database_manager_factory

import (
	"database/sql"
	"errors"
	"money-tracker/pkg/database/database_config"
	"money-tracker/pkg/database/database_credentials_formatter_factory"
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/database/database_type"
)

const (
	postgres              = "postgres"
	dbTypeNotSupportedMsg = "db type not supported"
)

type DatabaseManagerFactory struct {
}

func (dmf *DatabaseManagerFactory) ProduceDatabaseManager(databaseConfig *database_config.DatabaseConfig) (database_manager.DatabaseManagerInterface, error) {
	databaseCredentialsFormatterFactory := &database_credentials_formatter_factory.DatabaseCredentialsFormatterFactory{}

	switch *&databaseConfig.DatabaseType {
	case database_type.Postgres:
		dbConnection, err := sql.Open(postgres, databaseCredentialsFormatterFactory.ProduceFormattedCredentials(*databaseConfig))
		if err != nil {
			return nil, err
		}

		return &database_manager.PostgresDatabaseManager{DbConnection: dbConnection}, nil
	default:
		return nil, errors.New(dbTypeNotSupportedMsg)
	}
}
