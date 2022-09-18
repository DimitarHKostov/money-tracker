package register_validator

import (
	"money-tracker/pkg/database/database_manager"
	"money-tracker/types"
)

type DynamicRegisterValidator struct{}

func (drv *DynamicRegisterValidator) Validate(account types.Account, dbManager database_manager.DatabaseManagerInterface) (bool, error) {
	return dbManager.SelectAccountWithUsername(account.Username)
}
