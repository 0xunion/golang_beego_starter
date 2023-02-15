package controller

import (
	"errors"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/validation"
)

func ParseAndValidate[T any](item *T, controller beego.Controller) error {
	if err := controller.ParseForm(item); err != nil {
		return err
	}

	validate := validation.Validation{}
	if ok, _ := validate.Valid(item); !ok {
		if len(validate.Errors) > 0 {
			return validate.Errors[0]
		}
		return errors.New("validate error")
	}

	return nil
}
