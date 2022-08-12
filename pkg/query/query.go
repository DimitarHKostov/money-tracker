package query

const (
	registerAccountQuery         = `INSERT INTO Account(Email, Username, Password) VALUES ($1, $2, $3)`
	selectAccountsQuery          = `SELECT * FROM ACCOUNT`
	selectAccountByUsernameQuery = `SELECT Email, Username, Password FROM ACCOUNT WHERE Username = $1`
)

type Query struct{}

func (q *Query) GetRegisterAccountQuery() string {
	return registerAccountQuery
}

func (q *Query) GetSelectAccountsQuery() string {
	return selectAccountsQuery
}

func (q *Query) GetSelectAccountWithUsernameQuery() string {
	return selectAccountByUsernameQuery
}
