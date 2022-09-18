package register_validator

import "money-tracker/types"

type StaticRegisterValidatorInterface interface {
	Validate(types.Account) bool
}
