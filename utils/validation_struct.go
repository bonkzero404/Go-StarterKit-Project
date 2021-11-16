package utils

import (
	"go-boilerplate-clean-arch/domain/models"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func ValidateStruct(s interface{}) []*models.ErrorResponse {
	var errors []*models.ErrorResponse

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	err := validate.Struct(s)

	en_translations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse

			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Message = err.Translate(trans)
			errors = append(errors, &element)
		}
	}
	return errors
}
