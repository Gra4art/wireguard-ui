package router

import (
	"github.com/go-playground/validator/v10"
)

// NewValidator func
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validator struct
type Validator struct {
	validator *validator.Validate
}

// Validate func
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
