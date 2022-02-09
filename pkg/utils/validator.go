package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	validate := validator.New()

	_ = validate.RegisterValidation("uuid", func(v validator.FieldLevel) bool {
		field := v.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
