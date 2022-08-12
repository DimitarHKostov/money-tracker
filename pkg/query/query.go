package query

const (
	registerQuery = `INSERT INTO Account(Email, Username, Password) VALUES ($1, $2, $3)`
)

type Query struct{}

func (q *Query) GetRegisterQuery() string {
	return registerQuery
}
