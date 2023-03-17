package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/odanaraujo/crud-golang/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en2.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		_ = en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	}
	if errors.As(validation_err, &jsonValidationError) {
		var errorsCauses []rest_err.Causes
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("some fields are invalid", errorsCauses)
	}
	return rest_err.NewBadRequestError("error trying to convert fields")
}
