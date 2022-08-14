package database_manager

import (
	"database/sql"
	"money-tracker/pkg/query"
	"money-tracker/types"
)

type DatabaseManager struct {
	DbConnection *sql.DB
	Query        query.QueryInterface
}

func (dm *DatabaseManager) RegisterAccount(account types.Account) (sql.Result, error) {
	return dm.DbConnection.Exec(dm.Query.RegisterAccount(), account.Email, account.Username, account.Password)
}
