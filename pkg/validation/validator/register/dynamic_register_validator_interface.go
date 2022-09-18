package register_validator

import (
	"money-tracker/pkg/database/database_manager"
	"money-tracker/types"
)

type DynamicRegisterValidatorInterface interface {
	Validate(types.Account, database_manager.DatabaseManagerInterface) (bool, error)
}
