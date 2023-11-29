package util

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func FormatValidationErrors(err error) map[string]string {
	validationErrors := err.(validator.ValidationErrors)
	validationErrorMessages := make(map[string]string)

	for _, fieldError := range validationErrors {
		validationErrorMessages[fieldError.Field()] = fieldError.Tag()
	}

	return validationErrorMessages
}
