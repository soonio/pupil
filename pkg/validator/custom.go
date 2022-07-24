package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var phoneNumberRegexp = regexp.MustCompile(`^13\d{9}$|^14[5,7]\d{8}$|^15[^4]\d{8}$|^16\d{9}$|^17[0235678]\d{8}$|^18\d{9}$|^19\d{9}$`)

// isPhone 扩展一个验证手机号的方法, go-playground/validator自带的需要满足E146格式的手机号才行
func isPhone(fl validator.FieldLevel) bool {
	return phoneNumberRegexp.MatchString(fl.Field().String())
}

// YoN 是否验证
func YoN(fl validator.FieldLevel) bool {
	v := fl.Field().Int()
	return v == 0 || v == 1
}
