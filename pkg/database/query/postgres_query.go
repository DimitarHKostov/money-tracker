package query

type PostgresQuery struct{}

func (q *PostgresQuery) RegisterAccount() string {
	query := `INSERT INTO Account(Email, Username, Password) VALUES ($1, $2, $3)`
	return query
}

func (q *PostgresQuery) SelectAccounts() string {
	query := `SELECT * FROM ACCOUNT`
	return query
}

func (q *PostgresQuery) SelectAccountWithUsername() string {
	query := `SELECT Email, Username, Password FROM ACCOUNT WHERE Username = $1`
	return query
}

func (q *PostgresQuery) SelectAccountWithEmail() string {
	query := `SELECT Email, Username, Password FROM ACCOUNT WHERE Email = $1`
	return query
}
