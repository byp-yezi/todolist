package api

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"

	"todolist/pkg/ctl"
	"todolist/pkg/e"
	"todolist/pkg/util"
)

// func ErrorResponse(err error) *ctl.Response {
// 	if ve, ok := err.(validator.ValidationErrors); ok {
// 		for _, fieldError := range ve {
// 			field := conf.T(fmt.Sprintf("Field.%s", fieldError.Field()))
// 			tag := conf.T(fmt.Sprintf("Tag.%s", fieldError.Tag()))
// 			return ctl.RespSuccessWithData(fmt.Sprintf("%s%s", field, tag), e.InvalidParams)
// 		}
// 	}
// 	if _, ok := err.(*json.UnmarshalTypeError); ok {
// 		return ctl.RespError("JSON类型不匹配")
// 	}
// 	return ctl.RespError(err.Error(), e.ERROR)
// }

func ErrorResponse(err error) *ctl.Response {
	if errs, ok := err.(validator.ValidationErrors); ok {
		return ctl.RespError(util.RemoveTopStruct(errs.Translate(util.Trans)), e.InvalidParams)
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError("JSON类型不匹配")
	}
	return ctl.RespError(err.Error(), e.ERROR)
}
