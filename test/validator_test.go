package test

import (
	"EasyGo/middleware"
	"fmt"
	"testing"
)

func TestValidator(t *testing.T)  {
	req := middleware.RegisterReq{
		Username:       "Xargin",
		PasswordNew:    "ohno",
		PasswordRepeat: "ohn",
		Email:          "alex@abc.com",
	}

	err := middleware.ValidatorFunc(req)
	fmt.Println(err)
}
