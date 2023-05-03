package govalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()

	user := "user"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "password"
	confirmPassword := "password"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()

	user := "12345"

	err := validate.Var(user, "required,numeric")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()

	user := "10000"

	err := validate.Var(user, "required,numeric,min=5,max=10")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "user@email.com",
		Password: "password",
	}
	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "salah",
		Password: "password",
	}
	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.ActualTag(), "with error", fieldError.Error())
		}
	}
}

func TestStructCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}
	validate := validator.New()
	loginRequest := RegisterUser{
		Username:        "user@email.com",
		Password:        "password",
		ConfirmPassword: "12345",
	}
	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}
