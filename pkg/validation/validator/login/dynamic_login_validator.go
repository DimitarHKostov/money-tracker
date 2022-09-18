package login_validator

import (
	"money-tracker/pkg/database/database_manager"
	"money-tracker/types"
)

type DynamicLoginValidator struct{}

func (drv *DynamicLoginValidator) Validate(account types.Account, dbManager database_manager.DatabaseManagerInterface) (bool, error) {
	return true, nil
}
