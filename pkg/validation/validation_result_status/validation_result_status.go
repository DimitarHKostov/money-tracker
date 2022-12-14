package validation_result_status

type ValidationResultStatus int

const (
	Success ValidationResultStatus = iota
	Failure
)

func (vrs ValidationResultStatus) String() string {
	return [...]string{"Success", "Failure"}[vrs]
}
