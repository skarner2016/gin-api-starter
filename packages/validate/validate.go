package validate

import (
	"skarner2016/gin-api-starter/packages/translation"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func GetValidateError(form interface{}) (string, error) {
	validate = validator.New()
	if err := validate.Struct(form); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// 多语言
			msg := translation.GetValidateTransError(validate, err)
			return msg, err
		}
	}

	return "", nil
}

// func GetValidate() *validator.Validate {
// 	if validate != nil {
// 		validate = validator.New()
// 	}

// 	return validate
// }
