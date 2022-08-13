package validator

type ValidationResultMessage struct {
	Message string
}

func (vrm *ValidationResultMessage) GetMessage() string {
	return vrm.Message
}
