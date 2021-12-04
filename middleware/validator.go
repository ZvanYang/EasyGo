package middleware

import (
	"github.com/go-playground/validator/v10"
)

type RegisterReq struct {
	Username       string `validator:"gt=0"`
	PasswordNew    string `validator:"gt=0"`
	PasswordRepeat string `validator:"eqfield=PasswordNew"`
	Email          string `validator:"email"`
}

var validate = validator.New()

func ValidatorFunc(req RegisterReq) error {
	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
