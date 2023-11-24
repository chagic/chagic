package e

const (
	SUCCESS                     = 200
	UpdatePasswordSuccess       = 201
	NotExistInentifier          = 202
	ERROR                       = 500
	InvalidParams               = 400
	ErrorDatabase               = 40001
	ERROR_EXIST_USERNAME        = 40002
	ERROR_CREATE_USER_FAILED    = 40003
	ERROR_NOT_EXIST_USER        = 40004
	ERROR_PASSWORD_WRONG        = 40005
	ERROR_TOKEN_GENERATE_FAILED = 40006
	ERROR_GET_CHAT_USERS        = 40007

	WebsocketSuccessMessage = 50001
	WebsocketSuccess        = 50002
	WebsocketEnd            = 50003
	WebsocketOnlineReply    = 50004
	WebsocketOfflineReply   = 50005
	WebsocketLimit          = 50006
)
