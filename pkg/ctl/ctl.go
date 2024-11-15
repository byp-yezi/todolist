package ctl

import (
	"todolist/pkg/e"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

func RespSuccess(code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}
	r := &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}
	return r
}

func RespError(err error, data string, code ...int) *Response {
	status := e.ERROR
	if code != nil {
		status = code[0]
	}
	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
		Error:  err.Error(),
	}
	return r
}

func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}
	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
	return r
}

func RespList(item interface{}, total int64, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}
	r := &Response{
		Status: status,
		Data:   DataList{
			Item: item,
			Total: total,
		},
		Msg:    e.GetMsg(status),
	}
	return r
}
