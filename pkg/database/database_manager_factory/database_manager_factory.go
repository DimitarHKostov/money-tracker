package database_manager_factory

import (
	"database/sql"
	"errors"
	"fmt"
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/database/database_type"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test"
)

var (
	databaseConnectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
)

type DatabaseManagerFactory struct {
}

func (dmf *DatabaseManagerFactory) ProduceDatabaseManager(databaseType database_type.DatabaseType) (database_manager.DatabaseManagerInterface, error) {
	switch databaseType {
	case database_type.Postgres:
		dbConnection, err := sql.Open("postgres", databaseConnectionString)
		if err != nil {
			return nil, err
		}

		return &database_manager.PostgresDatabaseManager{DbConnection: dbConnection}, nil
	default:
		return nil, errors.New("db type not supported")
	}
}
