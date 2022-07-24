package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Error struct {
	errs validator.ValidationErrors
	uni  *ut.UniversalTranslator
	lang string
}

// Lang 设置使用的目标翻译语言
// 仅当目标语言有效时
func (e *Error) Lang(lang string) *Error {
	if _, found := e.uni.FindTranslator(lang); found {
		e.lang = lang
	}
	return e
}

// Error 实现error接口
func (e *Error) Error() string {
	var tips = make([]string, 0)
	t, _ := e.uni.GetTranslator(e.lang)
	for _, tip := range e.errs.Translate(t) {
		tips = append(tips, tip)
	}
	return strings.Join(tips, ", ")
}
