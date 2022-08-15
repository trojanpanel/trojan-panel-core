package api

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
)

var validate *validator.Validate

func InitValidator() {
	// Validate为单例对象
	validate = validator.New()
	// 注册校验对象
	_ = validate.RegisterValidation("validateStr", validateStr)
	_ = validate.RegisterValidation("validateEmail", validateEmail)
	_ = validate.RegisterValidation("validatePort", validatePort)
}

// 字符串必须是字母和数字的组合
func validateStr(f validator.FieldLevel) bool {
	field := f.Field().String()
	reg := "^[A-Za-z0-9]+$"
	compile := regexp.MustCompile(reg)
	return field == "" || compile.MatchString(field)
}

// 邮箱 只支持163 126 qq gmail
func validateEmail(f validator.FieldLevel) bool {
	if f.Field().Len() == 0 {
		return true
	}
	field := f.Field().String()
	reg := "^([A-Za-z0-9_\\-\\.])+\\@(163.com|126.com|qq.com|gmail.com)$"
	compile := regexp.MustCompile(reg)
	return field == "" || compile.MatchString(field)
}

// 正整数
func validatePositiveInt(f validator.FieldLevel) bool {
	field := strconv.FormatUint(f.Field().Uint(), 10)
	reg := "^[1-9]\\d*$"
	compile := regexp.MustCompile(reg)
	return compile.MatchString(field)
}

// 端口
func validatePort(f validator.FieldLevel) bool {
	field := strconv.FormatUint(f.Field().Uint(), 10)
	reg := "^([0-9]|[1-9]\\d{1,3}|[1-5]\\d{4}|6[0-4]\\d{4}|65[0-4]\\d{2}|655[0-2]\\d|6553[0-5])$"
	compile := regexp.MustCompile(reg)
	return compile.MatchString(field)
}
