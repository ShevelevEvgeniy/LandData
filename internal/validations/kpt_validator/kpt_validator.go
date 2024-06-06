package kpt_validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"regexp"
)

const (
	CadQuarter = "cad_quarter"
)

func ValidationKpt(validator *validator.Validate) error {
	err := validator.RegisterValidation(CadQuarter, validateCadQuarter)
	if err != nil {
		return errors.Wrap(err, "Failed registering kpt validation")
	}

	return nil
}

func validateCadQuarter(fl validator.FieldLevel) bool {
	cadQuarter := fl.Field().String()
	match, _ := regexp.MatchString(`^\d{2}:\d{2}:\d{7}$`, cadQuarter)

	return match
}
