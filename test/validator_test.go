package test

import (
	"EasyGo/validate"
	"fmt"
	"testing"
)

func TestValidator(t *testing.T)  {
	req := validate.RegisterReq{
		Username:       "Xargin",
		PasswordNew:    "ohno",
		PasswordRepeat: "ohn",
		Email:          "alex@abc.com",
	}

	err := validate.ValidatorFunc(req)
	fmt.Println(err)
}
