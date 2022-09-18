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

func (dm *PostgresDatabaseManager) SelectAccountWithUsername(username string) (bool, error) {
	var account types.Account
	err := dm.DbConnection.QueryRow(queryInstance.SelectAccountWithUsername(), username).Scan(&account)
	return account.Username == "", err
}
