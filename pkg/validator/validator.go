package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

type Validator struct {
	validator   *validator.Validate
	uni         *ut.UniversalTranslator
	defaultLang string
}

// Validate 验证方法
// 主要用于实现Echo#Validator，扩展自定义验证器
func (cv *Validator) Validate(i any) error {
	var e = cv.validator.Struct(i)
	if errs, ok := e.(validator.ValidationErrors); ok {
		return &Error{errs: errs, uni: cv.uni, lang: cv.defaultLang}
	}
	return e
}

func New(opt ...func(*Validator) *Validator) *Validator {
	v := new(Validator)
	for i := 0; i < len(opt); i++ {
		opt[i](v)
	}

	v.validator = validator.New()

	enLang := en.New()
	zhLang := zh.New()
	v.uni = ut.New(enLang, enLang, zhLang)

	var err error

	enTranslator, _ := v.uni.GetTranslator("en")
	err = entranslations.RegisterDefaultTranslations(v.validator, enTranslator)
	if err != nil {
		panic(err)
	}
	zhTranslator, _ := v.uni.GetTranslator("zh")
	err = zhtranslations.RegisterDefaultTranslations(v.validator, zhTranslator)
	if err != nil {
		panic(err)
	}

	_ = v.validator.RegisterValidation("phone", isPhone)
	_ = v.validator.RegisterValidation("yn", YoN)

	return v
}
