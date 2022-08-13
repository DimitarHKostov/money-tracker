package operation

type Operation int

const (
	Register Operation = iota
	Login
	Logout
	Refresh
	Calculate
)

func (o Operation) String() string {
	return [...]string{"Register", "Login", "Logout", "Refresh", "Calculate"}[o]
}
