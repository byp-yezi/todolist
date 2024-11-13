package e

var MsgFlags = map[int]string {
	SUCCESS: "操作成功",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}