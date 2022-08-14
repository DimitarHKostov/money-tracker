package validation_result_message

type ValidationResultMessage struct {
	Message string
}

func (vrm *ValidationResultMessage) GetMessage() string {
	return vrm.Message
}
