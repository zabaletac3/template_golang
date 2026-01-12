package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Errors []FieldError `json:"errors"`
}

func (v ValidationError) Error() string {
	var msgs []string
	for _, e := range v.Errors {
		msgs = append(msgs, fmt.Sprintf("%s: %s", e.Field, e.Message))
	}
	return strings.Join(msgs, ", ")
}

func Validate(err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		fields := make([]FieldError, len(ve))
		for i, fe := range ve {
			fields[i] = FieldError{
				Field:   toSnakeCase(fe.Field()),
				Message: msgForTag(fe),
			}
		}
		return ValidationError{Errors: fields}
	}
	return err
}

func msgForTag(fe validator.FieldError) string {
	field := toSnakeCase(fe.Field())

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return "must be a valid email"
	case "min":
		return fmt.Sprintf("must be at least %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", fe.Param())
	case "gte":
		return fmt.Sprintf("must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("must be less than or equal to %s", fe.Param())
	default:
		return fmt.Sprintf("failed on %s validation", fe.Tag())
	}
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteByte('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v
	}
}
