package Visitors

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-starter/Models"
	"go-starter/Validation"
)

func RegisterValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"username": validation.Validate(
			user.Username,
			Validation.RequiredRule(),
			Validation.MinMaxRule(2, 20)),
		"email": validation.Validate(
			user.Email,
			Validation.RequiredRule(),
			Validation.IsEmailRule(),
			Validation.MinMaxRule(5, 35)),
		"password": validation.Validate(
			user.Password,
			Validation.RequiredRule(),
			Validation.MinMaxRule(8, 15)),
	}
}

func LoginValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"email": validation.Validate(
			user.Email,
			Validation.RequiredRule(),
			Validation.IsEmailRule(),
			Validation.MinMaxRule(5, 35)),
		"password": validation.Validate(
			user.Password,
			Validation.RequiredRule(),
			Validation.MinMaxRule(8, 15)),
	}
}
