package database_type

type DatabaseType int

const (
	Postgres DatabaseType = iota
)

func (dt DatabaseType) String() string {
	return [...]string{"Postgres"}[dt]
}
