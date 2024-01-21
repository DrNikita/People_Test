package validation

import (
	"fmt"
	"github.com/DrNikita/People/internal/model"
	"github.com/go-playground/validator/v10"
)

func ValidatePerson(person model.SupplementedPerson) (validationErr string, err error) {
	validate, translator, err := GetValidation()
	if err != nil {
		return "", err
	}
	err = validate.Struct(person)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fieldName := e.Namespace()
			validationErr += fmt.Sprintf("%s - %s; ", fieldName, e.Translate(translator))
		}
		return validationErr, nil
	}
	return "", nil
}
