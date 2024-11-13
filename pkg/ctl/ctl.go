package ctl

import "todolist/pkg/e"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func RespSuccess(code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}
	r := &Response{
		Status: status,
		Data: "操作成功",
		Msg: e.GetMsg(status),
	}
	return r
}
