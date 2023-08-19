package infra

import (
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	uni *ut.UniversalTranslator
)

func ValidateStruct(data any) []string {

	validate := validator.New()
	err := validate.Struct(data)
	var errors []string

	pt := pt.New()

	uni = ut.New(pt, pt)

	trans, _ := uni.GetTranslator("pt_BR")

	pt_BR.RegisterDefaultTranslations(validate, trans)

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Translate(trans))
		}
	}
	return errors
}
