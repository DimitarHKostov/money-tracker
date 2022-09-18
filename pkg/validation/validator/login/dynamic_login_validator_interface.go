package login_validator

import (
	"money-tracker/pkg/database/database_manager"
	"money-tracker/types"
)

type DynamicLoginValidatorInterface interface {
	Validate(types.Account, database_manager.DatabaseManagerInterface) (bool, error)
}
