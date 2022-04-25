package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

func NewValidator() CustomValidator {
	cv := CustomValidator{}
	en := en.New()
	uni := ut.New(en, en)
	cv.Trans, _ = uni.GetTranslator("en")
	cv.Validate = validator.New()

	cv.Validate.RegisterTagNameFunc(getTagName)
	en_translations.RegisterDefaultTranslations(cv.Validate, cv.Trans)
	return cv
}

func ValidateStruct(s interface{}) (map[string][]string, bool) {
	cv := NewValidator()
	err := cv.Validate.Struct(s)

	values := map[string][]string{}

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			field := e.Field()
			message := e.Translate(cv.Trans)
			values[field] = append(values[field], message)
		}

		return values, false
	}

	return values, true
}

func getTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}
