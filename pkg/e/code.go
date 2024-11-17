package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	// 用户
	ErrorExistUser               = 10002
	ErrorNotExistUser            = 10003
	ErrorUserOrPasswordIncorrect = 10004

	// 认证
	ErrorAuthTokenFail = 20001
	ErrorAuthNotFound  = 20003

	// 任务
	ErrorTaskNotFound        = 30001
	ErrorTaskStatusParameter = 30002
	ErrorTaskNotFoundByUser  = 30003
)
