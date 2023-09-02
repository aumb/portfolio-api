package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func Validate(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate.RegisterTagNameFunc(jsonTagNameFunc)

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ParseValidationErrors(errs []*ErrorResponse) string {
	errMsgs := make([]string, 0)

	for _, err := range errs {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}

	return strings.Join(errMsgs, " and ")
}

func jsonTagNameFunc(field reflect.StructField) string {
	name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	// skip if tag key says it should be ignored
	if name == "-" {
		return ""
	}
	return name
}
