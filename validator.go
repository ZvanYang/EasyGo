package main

import (
	"github.com/go-playground/validator/v10"
)

type registerReq struct {
	Username       string `validator:"gt=0"`
	PasswordNew    string `validator:"gt=0"`
	PasswordRepeat string `validator:"eqfield=PasswordNew"`
	Email          string `validator:"email"`
}

var validate = validator.New()

func validatorFunc(req registerReq) error {
	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
