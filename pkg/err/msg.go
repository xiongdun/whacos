package err

var MsgFlags = map[int]string{
	Success:                    "请求成功！",
	Error:                      "请求失败！",
	InvalidParams:              "请求参数错误",
	ErrorNotExistData:          "数据不存在！",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
