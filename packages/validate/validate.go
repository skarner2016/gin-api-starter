package validate

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func GetValidateError(form interface{}) (string, error) {
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// 字段
			// fmt.Println(err.Field())
			//
			// fmt.Println(err.Tag())

			// TODO: 多语言
			field := err.Field()
			tag := err.Tag()
			msg := fmt.Sprintf("%s-%s", field, tag)

			return msg, errors.New("")
		}
	}

	return "", nil
}

func GetValidate() *validator.Validate {
	if validate != nil {
		validate = validator.New()
	}

	return validate
}
