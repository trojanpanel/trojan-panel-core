package api

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	// Validate为单例对象
	validate = validator.New()
}
