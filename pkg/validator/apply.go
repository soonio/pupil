package validator

// 配置项

// WithDefaultLanguage 设置默认配置
// lang 当前可选值 zh | en
func WithDefaultLanguage(lang string) func(v *Validator) *Validator {
	return func(v *Validator) *Validator {
		v.defaultLang = lang
		return v
	}
}
