package api

import (
	"encoding/json"
	"fmt"
	conf "todolist/config"
	"todolist/pkg/ctl"
	"todolist/pkg/e"

	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) *ctl.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", fieldError.Field()))
			tag := conf.T(fmt.Sprintf("Tag.%s", fieldError.Tag()))
			return ctl.RespError(err, fmt.Sprintf("%s%s", field, tag))
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(err, "JSON类型不匹配")
	}
	return ctl.RespError(err, err.Error(), e.ERROR)
}