package database_manager

import (
	"database/sql"
	"money-tracker/pkg/database/query"
	"money-tracker/types"
)

var (
	queryInstance = query.PostgresQuery{}
)

type PostgresDatabaseManager struct {
	DbConnection *sql.DB
}

func (dm *PostgresDatabaseManager) RegisterAccount(account types.Account) (sql.Result, error) {
	return dm.DbConnection.Exec(queryInstance.RegisterAccount(), account.Email, account.Username, account.Password)
}
