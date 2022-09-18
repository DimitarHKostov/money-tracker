package validator_factory

import (
	"money-tracker/pkg/operation"
	"money-tracker/pkg/validation/validator"
	lv "money-tracker/pkg/validation/validator/login"
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
	case operation.Login:
		return &lv.LoginValidator{
			DynamicLoginValidator: &lv.DynamicLoginValidator{},
		}
	default:
		return nil
	}

}
