package validator

import (
	"fmt"
	"strings"

	playground "github.com/go-playground/validator/v10"
)

var validate = playground.New()

func Validate(v any) error {
	return validate.Struct(v)
}

func ParseErrors(err error) []ValidationError {

	validationErrors, ok := err.(playground.ValidationErrors)
	if !ok {
		return nil
	}

	errors := make([]ValidationError, 0, len(validationErrors))

	for _, field := range validationErrors {

		errors = append(errors, ValidationError{
			Field:   strings.ToLower(field.Field()),
			Message: message(field),
		})

	}

	return errors
}

func message(field playground.FieldError) string {

	switch field.Tag() {

	case "required":
		return "is required"

	case "gt":
		return fmt.Sprintf("must be greater than %s", field.Param())

	case "gte":
		return fmt.Sprintf("must be greater than or equal to %s", field.Param())

	case "max":
		return fmt.Sprintf("must not exceed %s characters", field.Param())

	default:
		return field.Error()
	}
}
