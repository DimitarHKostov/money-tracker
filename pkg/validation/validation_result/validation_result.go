package validation_result

import (
	"money-tracker/pkg/validation/validation_result_message"
	"money-tracker/pkg/validation/validation_result_status"
)

type ValidationResult struct {
	ValidationResultStatus  validation_result_status.ValidationResultStatus
	ValidationResultMessage validation_result_message.ValidationResultMessageInterface
}
