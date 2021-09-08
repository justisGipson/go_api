package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	// new validator for lesson model
	validate := validator.New()

	// custom validation for uuid.UUID fields
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})
	return validate
}

func ValidatorErrors(err error) map[string]string {
	// define fields map
	fields := map[string]string{}

	// make error msg for each invalid field
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}
	return fields
}
