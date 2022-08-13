package query

const (
	registerAccount         = `INSERT INTO Account(Email, Username, Password) VALUES ($1, $2, $3)`
	selectAccounts          = `SELECT * FROM ACCOUNT`
	selectAccountByUsername = `SELECT Email, Username, Password FROM ACCOUNT WHERE Username = $1`
	selectAccountByEmail    = `SELECT Email, Username, Password FROM ACCOUNT WHERE Email = $1`
)

type Query struct{}

func (q *Query) RegisterAccount() string {
	return registerAccount
}

func (q *Query) SelectAccounts() string {
	return selectAccounts
}

func (q *Query) SelectAccountWithUsername() string {
	return selectAccountByUsername
}

func (q *Query) SelectAccountWithEmail() string {
	return selectAccountByEmail
}
