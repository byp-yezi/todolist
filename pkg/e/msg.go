package e

var MsgFlags = map[int]string{
	SUCCESS:       "操作成功",
	ERROR:         "操作失败",
	InvalidParams: "请求参数错误",

	ErrorExistUser:               "已存在该用户名",
	ErrorNotExistUser:            "该用户不存在",
	ErrorUserOrPasswordIncorrect: "账号或密码不正确",

	ErrorAuthTokenFail:    "Token认证失败",
	ErrorAuthTokenExpired: "Token已失效",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
