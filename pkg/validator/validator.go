package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

var phoneNumberRegexp *regexp.Regexp

func (cv *Validator) Validate(i any) error {
	return cv.validator.Struct(i)
}

// isPhone 扩展一个验证手机号的方法, go-playground/validator自带的需要满足E146格式的手机号才行
func isPhone(fl validator.FieldLevel) bool {
	return phoneNumberRegexp.MatchString(fl.Field().String())
}

// YoN 是否验证
func YoN(fl validator.FieldLevel) bool {
	v := fl.Field().Int()
	return v == 0 || v == 1
}

// 只支持数字和字符串
func requiredIn(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), " ")

	field, kind, _, found := fl.GetStructFieldOKAdvanced2(fl.Parent(), params[0])
	if !found {
		return false
	}

	var sv string

	switch kind {
	case reflect.String:
		sv = field.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		sv = fmt.Sprintf("%d", field.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		sv = fmt.Sprintf("%d", field.Uint())
	case reflect.Float32, reflect.Float64:
		sv = fmt.Sprintf("%f", field.Float())
	default:
		return false
	}

	var in bool
	for _, need := range params[1:] {
		if need == sv {
			in = true
		}
	}
	if in {
		field, kind, _, found = fl.GetStructFieldOKAdvanced2(fl.Parent(), params[0])
		if !found {
			return false
		}
		switch kind {
		case reflect.String:
			return field.String() != ""
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return field.Int() != 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return field.Uint() != 0
		case reflect.Float32, reflect.Float64:
			return field.Float() != 0
		}
	}
	return true
}

func New() *Validator {
	v := &Validator{validator: validator.New()}

	phoneNumberRegexp = regexp.MustCompile(`^13[\d]{9}$|^14[5,7]\d{8}$|^15[^4]\d{8}$|^16[\d]{9}$|^17[0235678]\d{8}$|^18[\d]{9}$|^19[\d]{9}$`)
	_ = v.validator.RegisterValidation("phone", isPhone)
	_ = v.validator.RegisterValidation("yn", YoN)
	_ = v.validator.RegisterValidation("required_in", requiredIn)
	return v
}
