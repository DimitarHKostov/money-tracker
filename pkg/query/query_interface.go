package query

type QueryInterface interface {
	GetRegisterAccountQuery() string
	GetSelectAccountsQuery() string
	GetSelectAccountWithUsernameQuery() string
}
