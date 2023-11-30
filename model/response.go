package model

type Response struct {
	ResponseCode string      `json:"responseCode" swagger:"description(ResponseCode)"`
	ResponseDesc string      `json:"responseDesc" swagger:"description(ResponseDesc)"`
	Data         interface{} `json:"data,omitempty" swagger:"description(Data)"`
}
