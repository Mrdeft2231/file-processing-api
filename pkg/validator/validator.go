package validator

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	// Register custom validation function for username
	_ = validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		username := fl.Field().String()

		// Username must be between 3 and 32 characters
		if len(username) < 3 || len(username) > 32 {
			return false
		}

		// Username can only contain letters, numbers, underscores and hyphens
		for _, char := range username {
			if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '_' && char != '-' {
				return false
			}
		}

		return true
	})
}

func ValidateRequest(req interface{}) error {
	if err := validate.Struct(req); err != nil {
		return errors.New(parseValidationError(err))
	}
	return nil
}

func parseValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				messages = append(messages, fmt.Sprintf("Field %s is required", e.Field()))
			case "username":
				messages = append(messages, fmt.Sprintf("Field %s must be a valid username", e.Field()))
			case "min":
				messages = append(messages, fmt.Sprintf("Field %s must have at least %s characters", e.Field(), e.Param()))
			case "max":
				messages = append(messages, fmt.Sprintf("Field %s must have no more than %s characters", e.Field(), e.Param()))
			default:
				messages = append(messages, fmt.Sprintf("Field %s is invalid", e.Field()))
			}
		}
		return strings.Join(messages, "; ")
	}
	return "Validation error not found"
}
