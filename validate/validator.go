package validate

import (
	"github.com/go-playground/validator/v10"
)

type RegisterReq struct {
	Username       string `validate:"gt=0"`
	PasswordNew    string `validate:"gt=0"`
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	Email          string `validate:"email"`
}

var validate = validator.New()

func ValidatorFunc(req RegisterReq) error {
	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
