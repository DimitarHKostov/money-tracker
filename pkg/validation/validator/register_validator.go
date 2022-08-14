package validator

import (
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/validation/validation_result"
	"money-tracker/pkg/validation/validation_result_message"
	"money-tracker/pkg/validation/validation_result_status"
	"net/http"
)

type RegisterValidator struct{}

func (rsv RegisterValidator) Validate(*http.Request, database_manager.DatabaseManagerInterface) *validation_result.ValidationResult {
	return &validation_result.ValidationResult{
		ValidationResultStatus: validation_result_status.Success,
		ValidationResultMessage: &validation_result_message.ValidationResultMessage{
			Message: validation_result_status.Success.String(),
		},
	}
}
