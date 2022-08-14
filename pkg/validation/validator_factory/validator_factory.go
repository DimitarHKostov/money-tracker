package validator_factory

import (
	"money-tracker/pkg/operation"
	"money-tracker/pkg/validation/validator"
)

type ValidatorFactory struct {
}

func (vf *ValidatorFactory) ProduceValidator(op operation.Operation) validator.ValidatorInterface {
	switch op {
	case operation.Register:
		return &validator.RegisterSubValidator{}
	default:
		return nil
	}

}
