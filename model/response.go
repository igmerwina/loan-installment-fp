package model

type Response struct {
	RespCode int         `json:"status" form:"status"`
	Message  string      `json:"message" form:"message"`
	Data     interface{} `json"data" form:"data"`
}
