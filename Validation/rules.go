package Validation

import (
	"github.com/bykovme/gotrans"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func RequiredRule() validation.Rule {
	return validation.Required.Error(gotrans.T("required"))
}

func MinMaxRule(min int, max int) validation.Rule {
	return validation.Length(min, max).Error(gotrans.T("min_max"))
}

func IsEmailRule() validation.Rule {
	return is.Email.Error(gotrans.T("email"))
}
