package query

const (
	registerAccount         = `INSERT INTO Account(Email, Username, Password) VALUES ($1, $2, $3)`
	selectAccounts          = `SELECT * FROM ACCOUNT`
	selectAccountByUsername = `SELECT Email, Username, Password FROM ACCOUNT WHERE Username = $1`
	selectAccountByEmail    = `SELECT Email, Username, Password FROM ACCOUNT WHERE Email = $1`
)

type PostgresQuery struct{}

func (q *PostgresQuery) RegisterAccount() string {
	return registerAccount
}

func (q *PostgresQuery) SelectAccounts() string {
	return selectAccounts
}

func (q *PostgresQuery) SelectAccountWithUsername() string {
	return selectAccountByUsername
}

func (q *PostgresQuery) SelectAccountWithEmail() string {
	return selectAccountByEmail
}
