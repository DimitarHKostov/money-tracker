package validator

import "errors"

type Validator struct{}

func (v *Validator) ValidateUsername(username string) error {
	if len(username) < 6 {
		return errors.New("Not enough symbols")
	}

	hasUppercaseLetter := false
	hasLowercaseLetter := false
	hasSpecialSymbol := false
	hasNumber := false

	for _, letter := range username {
		if letter >= 'a' && letter <= 'z' {
			hasLowercaseLetter = true
		}

		if letter >= 'A' && letter <= 'Z' {
			hasUppercaseLetter = true
		}

		if letter >= '0' && letter <= '9' {
			hasNumber = true
		}

		if letter >= '!' && letter <= '/' {
			hasSpecialSymbol = true
		}
	}

	if !hasUppercaseLetter || !hasLowercaseLetter || !hasSpecialSymbol || !hasNumber {
		return errors.New("validation failed")
	}

	return nil
}
