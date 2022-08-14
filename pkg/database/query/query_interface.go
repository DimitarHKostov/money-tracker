package query

type QueryInterface interface {
	RegisterAccount() string
	SelectAccounts() string
	SelectAccountWithUsername() string
	SelectAccountWithEmail() string
}
