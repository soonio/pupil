package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

type Error struct {
	errs validator.ValidationErrors
	uni  *ut.UniversalTranslator
	lang string
}

func (e *Error) Lang(lang string) error {
	e.lang = lang
	return e
}

func (e *Error) Error() string {
	var tips = make([]string, 0)

	var t ut.Translator
	var found bool
	if t, found = e.uni.FindTranslator(e.lang); !found {
		t, _ = e.uni.GetTranslator("en")
	}

	for _, tip := range e.errs.Translate(t) {
		tips = append(tips, tip)
	}

	return strings.Join(tips, ", ")
}

type Validator struct {
	validator *validator.Validate
	uni       *ut.UniversalTranslator
}

// Validate 验证方法
// 主要用于实现Echo#Validator，扩展自定义验证器
func (cv *Validator) Validate(i any) error {
	var e = cv.validator.Struct(i)
	if errs, ok := e.(validator.ValidationErrors); ok {
		return &Error{errs: errs, uni: cv.uni, lang: "en"}
	}
	return e
}

func New() *Validator {
	v := new(Validator)
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
