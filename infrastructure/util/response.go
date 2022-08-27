package util

import "encoding/json"

// Response 回复消息
type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewResponse(t string, data interface{}) *Response {
	return &Response{Type: t, Data: data}
}

// Json 消息转json字符串
func (r *Response) Json() []byte {
	if bytes, err := json.Marshal(r); err != nil {
		return []byte("")
	} else {
		return bytes
	}
}
