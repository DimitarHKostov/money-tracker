package database_manager

import (
	"database/sql"
	"money-tracker/types"
)

type DatabaseManagerInterface interface {
	RegisterAccount(types.Account) (sql.Result, error)
}
