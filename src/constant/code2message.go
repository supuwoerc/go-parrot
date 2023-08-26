package constant

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 10000
	InvalidToken  = 10001
)

var code2message = map[int]string{
	SUCCESS:       "操作成功",
	ERROR:         "操作失败",
	InvalidParams: "参数校验失败",
	InvalidToken:  "鉴权失败",
}

func GetMessage(code int) string {
	message, exist := code2message[code]
	if exist {
		return message
	}
	return code2message[ERROR]
}
