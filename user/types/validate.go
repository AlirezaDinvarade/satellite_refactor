package types

import "github.com/go-playground/validator/v10"

type CustomValidate struct {
	*validator.Validate
}

func NewValidator() *CustomValidate {
	v := validator.New()
	validatorWrapper := &CustomValidate{Validate: v}

	validatorWrapper.RegisterValidation("startswith09", validatorWrapper.startswith09)

	return validatorWrapper
}

func (v *CustomValidate) startswith09(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	return len(phone) == 11 && phone[:2] == "09"
}
