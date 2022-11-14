package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type MyStruct struct {
	String string `validate:"email"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()

	s := MyStruct{String: "yo@domail.com"}

	err := validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	// s.String = "not awesome"
	// err = validate.Struct(s)
	// if err != nil {
	// 	fmt.Printf("Err(s):\n%+v\n", err)
	// }
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
