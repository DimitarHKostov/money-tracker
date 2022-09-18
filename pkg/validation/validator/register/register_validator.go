package register_validator

import (
	"encoding/json"
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/validation/validation_result"
	"money-tracker/pkg/validation/validation_result_message"
	"money-tracker/pkg/validation/validation_result_status"
	"money-tracker/types"
)

type RegisterValidator struct {
	StaticRegisterValidator  StaticRegisterValidatorInterface
	DynamicRegisterValidator DynamicRegisterValidatorInterface
}

func (rsv *RegisterValidator) Validate(bodyBytes []byte, dbManager database_manager.DatabaseManagerInterface) *validation_result.ValidationResult {
	var account types.Account
	err := json.Unmarshal(bodyBytes, &account)
	if err != nil {
		return rsv.returnFailure(err)
	}

	return rsv.returnSuccess()
}

func (rsv *RegisterValidator) returnSuccess() *validation_result.ValidationResult {
	return &validation_result.ValidationResult{
		ValidationResultStatus: validation_result_status.Success,
		ValidationResultMessage: &validation_result_message.ValidationResultMessage{
			Message: validation_result_status.Success.String(),
		},
	}
}

func (rsv *RegisterValidator) returnFailure(err error) *validation_result.ValidationResult {
	return &validation_result.ValidationResult{
		ValidationResultStatus: validation_result_status.Failure,
		ValidationResultMessage: &validation_result_message.ValidationResultMessage{
			Message: err.Error(),
		},
	}
}

func (rsv *RegisterValidator) checkForCollision(account types.Account) (bool, error) {
	return false, nil
}
