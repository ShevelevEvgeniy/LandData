package land_plots_validatior

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

const (
	CadNumber = "cad_number"
)

func ValidationLandPlots(validator *validator.Validate) error {
	err := validator.RegisterValidation(CadNumber, validateCadNumber)
	if err != nil {
		return errors.Wrap(err, "Failed registering kpt validation")
	}

	return nil
}

func validateCadNumber(fl validator.FieldLevel) bool {
	cadNumber := fl.Field().String()
	match, _ := regexp.MatchString(`^\d{2}:\d{2}:\d{7}:\d{1,5}$`, cadNumber)

	return match
}
