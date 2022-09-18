package register_validator

import "money-tracker/types"

type StaticRegisterValidator struct {
}

func (srv *StaticRegisterValidator) Validate(account types.Account) bool {
	if result := srv.isUsernameValid(account.Username) && srv.isEmailValid(account.Email); result == false {
		return true
	}

	return false
}

func (srv *StaticRegisterValidator) isEmailValid(email string) bool {
	return true
}

func (srv *StaticRegisterValidator) isUsernameValid(username string) bool {
	if len(username) < 6 {
		return false
	}

	hasLowercaseSymbol := false
	hasUppercaseSymbol := false
	hasSpecialSymbol := false
	hasNumber := false

	for _, symbol := range username {
		if srv.isNumber(symbol) {
			hasNumber = true
		}

		if srv.isLowercaseLetter(symbol) {
			hasLowercaseSymbol = true
		}

		if srv.isUppercaseLetter(symbol) {
			hasUppercaseSymbol = true
		}

		if srv.isSpecialSymbol(symbol) {
			hasSpecialSymbol = true
		}
	}

	return hasLowercaseSymbol && hasUppercaseSymbol && hasSpecialSymbol && hasNumber
}

func (srv *StaticRegisterValidator) isNumber(symbol rune) bool {
	return symbol >= '0' && symbol <= '9'
}

func (srv *StaticRegisterValidator) isLowercaseLetter(symbol rune) bool {
	return symbol >= 'a' && symbol <= 'z'
}

func (srv *StaticRegisterValidator) isUppercaseLetter(symbol rune) bool {
	return symbol >= 'A' && symbol <= 'Z'
}

func (srv *StaticRegisterValidator) isSpecialSymbol(symbol rune) bool {
	return symbol >= '!' && symbol <= '/'
}
