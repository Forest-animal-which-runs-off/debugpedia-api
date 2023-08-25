package validator

import (
	"debugpedia-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IDebugValidator interface {
	DebugValidate(debug model.Debug) error
}

type debugValidator struct{}

func NewDebugValidator() IDebugValidator {
	return &debugValidator{}
}

func (tv *debugValidator) DebugValidate(Debug model.Debug) error {
	return validation.ValidateStruct(&Debug,
		validation.Field(
			&Debug.Title,
			validation.Required.Error("title is required"),
		),
		validation.Field(
			&Debug.Body,
			validation.Required.Error("body is required"),
		),
		validation.Field(
			&Debug.Techs,
			validation.Required.Error("at least one tech is required"),
		),
	)
}