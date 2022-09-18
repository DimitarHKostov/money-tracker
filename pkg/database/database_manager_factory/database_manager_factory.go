package database_manager_factory

import (
	"database/sql"
	"errors"
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/database/database_type"
)

const (
	postgres              = "postgres"
	dbTypeNotSupportedMsg = "db type not supported"
)

type DatabaseManagerFactory struct {
}

func (dmf *DatabaseManagerFactory) ProduceDatabaseManager(databaseType database_type.DatabaseType, dbConnectionString string) (database_manager.DatabaseManagerInterface, error) {
	switch databaseType {
	case database_type.Postgres:
		dbConnection, err := sql.Open(postgres, dbConnectionString)
		if err != nil {
			return nil, err
		}

		return &database_manager.PostgresDatabaseManager{DbConnection: dbConnection}, nil
	default:
		return nil, errors.New(dbTypeNotSupportedMsg)
	}
}
