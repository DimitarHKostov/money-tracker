package validator

type ValidatorInterface interface {
	ValidateUsername(string) error
}
