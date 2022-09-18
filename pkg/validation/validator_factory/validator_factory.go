package validator_factory

import (
	"money-tracker/pkg/operation"
	"money-tracker/pkg/validation/validator"
	rv "money-tracker/pkg/validation/validator/register"
)

type ValidatorFactory struct {
}

func (vf *ValidatorFactory) ProduceValidator(op operation.Operation) validator.ValidatorInterface {
	switch op {
	case operation.Register:
		return &rv.RegisterValidator{
			StaticRegisterValidator:  &rv.StaticRegisterValidator{},
			DynamicRegisterValidator: &rv.DynamicRegisterValidator{},
		}
	default:
		return nil
	}

}
