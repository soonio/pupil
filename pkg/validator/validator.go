package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

type Error []string

func (e Error) Error() string {
	return strings.Join(e, ", ")
}

type Validator struct {
	validator  *validator.Validate
	translator ut.Translator
}

// Validate 验证方法
// 主要用于实现Echo#Validator，扩展echo的方法
func (cv *Validator) Validate(i any) error {
	var e = cv.validator.Struct(i)
	if errs, ok := e.(validator.ValidationErrors); ok {
		tips := make(Error, 0)
		for _, tip := range errs.Translate(cv.translator) {
			tips = append(tips, tip)
		}
		return tips
	}
	return e
}

func New() *Validator {
	v := new(Validator)

	lang := en.New()
	uni := ut.New(lang, lang)

	v.translator, _ = uni.GetTranslator("en")

	v.validator = validator.New()
	err := translations.RegisterDefaultTranslations(v.validator, v.translator)
	if err != nil {
		panic(err)
	}

	_ = v.validator.RegisterValidation("phone", isPhone)
	_ = v.validator.RegisterValidation("yn", YoN)

	return v
}
