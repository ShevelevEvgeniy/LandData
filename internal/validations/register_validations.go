package validations

import (
	landPlotsValidatior "github.com/ShevelevEvgeniy/app/internal/validations/land_plots_validatior"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func RegisterValidations(validate *validator.Validate) error {
	listRegistered := []func(*validator.Validate) error{
		landPlotsValidatior.ValidationLandPlots,
	}

	for _, fn := range listRegistered {
		err := fn(validate)
		if err != nil {
			return errors.Wrap(err, "Failed registering validation")
		}
	}

	return nil
}
