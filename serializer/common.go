package serializer

import "chagic/pkg/e"

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func ResultOK(data interface{}) Response {
	code := e.SUCCESS
	return Response{
		Code:    code,
		Msg:     e.GetMsg(code),
		Success: true,
		Data:    data,
	}
}

func ResultErr(code int) Response {
	return Response{
		Code:    code,
		Msg:     e.GetMsg(code),
		Success: false,
	}
}
