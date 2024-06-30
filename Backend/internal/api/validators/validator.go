package validators

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ExpiryDateValidation is a custom validator function that checks if a field's value matches the "MM/YY" format.
// It uses a regular expression to perform the validation.
func ExpiryDateValidation(fl validator.FieldLevel) bool {
	expiryDate := fl.Field().String()
	match, _ := regexp.MatchString(`^(0[1-9]|1[0-2])\/([0-9]{2})$`, expiryDate)
	return match
}

// TranslateValidationErrors translates validation errors into a slice of readable error messages.
// It takes an error object returned by the validator and returns a slice of strings with human-readable error messages.
func TranslateValidationErrors(err error) []string {
	var errMsg []string

	// Loop through each validation error and append a formatted message to the errMsg slice
	for _, err := range err.(validator.ValidationErrors) {
		errMsg = append(errMsg, fmt.Sprintf("%s %s", err.Field(), getValidationErrorMessage(err)))
	}

	return errMsg
}

// getValidationErrorMessage returns a human-readable error message based on the validation tag of the field error.
// It handles different validation tags such as "required", "alpha", "credit_card", "expirydate", "gt", "len", and "numeric".
func getValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "is required"
	case "alpha":
		return "must only contain alphabetic characters"
	case "credit_card":
		return "must be a valid credit card number"
	case "expirydate":
		return "must be in MM/YY format"
	case "gt":
		return fmt.Sprintf("must be greater than %s", err.Param())
	case "len":
		return fmt.Sprintf("must be exactly %s characters", err.Param())
	case "numeric":
		return "must be numeric"
	default:
		return "is invalid"
	}
}
