package translation

import (
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	local string
	uni   *ut.UniversalTranslator
	tans  ut.Translator
)

func Setup() {
	zh := zh2.New()
	en := en2.New()
	uni = ut.New(en, zh)
	local = "en"

	tans, _ = uni.GetTranslator(local)
}

func GetValidateTransError(validate *validator.Validate, err validator.FieldError) string {
	if uni == nil {
		Setup()
	}

	switch local {
	case "en":
		en_translations.RegisterDefaultTranslations(validate, tans)
	case "zh":
		zh_translations.RegisterDefaultTranslations(validate, tans)
	default:
		en_translations.RegisterDefaultTranslations(validate, tans)
	}

	return err.Translate(tans)
}
