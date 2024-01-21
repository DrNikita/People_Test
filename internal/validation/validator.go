package validation

import (
	"github.com/go-playground/locales/ru"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func GetValidation() (validate *validator.Validate, translator ut.Translator, err error) {
	validate = validator.New()
	ruLocale := ru.New()
	uni := ut.New(ruLocale, ruLocale)

	translator, _ = uni.GetTranslator("ru")

	err = validate.RegisterTranslation(
		"required",
		translator,
		func(ut ut.Translator) error {
			return ut.Add("required", "обязательное поле", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)

	return
}
