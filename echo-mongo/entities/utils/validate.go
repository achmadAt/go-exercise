package utils

import (
	validator "github.com/go-playground/validator/v10"
)

func Validate(item interface{}) error {
	validate := validator.New()
	if err := validate.Struct(item); err != nil {
		return err
	}
	return nil
}
