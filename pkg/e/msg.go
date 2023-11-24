package e

var MsgFlags = map[int]string{
	SUCCESS:                     "ok",
	ERROR_EXIST_USERNAME:        "用户已存在",
	ERROR_CREATE_USER_FAILED:    "注册用户失败",
	ERROR_NOT_EXIST_USER:        "用户不存在",
	ERROR_PASSWORD_WRONG:        "密码错误",
	ERROR_TOKEN_GENERATE_FAILED: "token生成失败",
	UpdatePasswordSuccess:       "修改密码成功",
	NotExistInentifier:          "该第三方账号未绑定",
	ERROR:                       "fail",
	InvalidParams:               "请求参数错误",
	ErrorDatabase:               "数据库操作出错,请重试",
	WebsocketSuccessMessage:     "解析content内容信息",
	WebsocketSuccess:            "发送信息，请求历史纪录操作成功",
	WebsocketEnd:                "请求历史纪录，但没有更多记录了",
	WebsocketOnlineReply:        "针对回复信息在线应答成功",
	WebsocketOfflineReply:       "针对回复信息离线回答成功",
	WebsocketLimit:              "请求收到限制",
	ERROR_GET_CHAT_USERS:        "获取聊天用户失败",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
