package ws

import "encoding/json"

// Message 消息
type Message struct {
	Type int
	Data []byte
}

func NewMessage(mType int, data []byte) *Message {
	return &Message{Type: mType, Data: data}
}

// Parse2Command 将消息解析成指令
func (m *Message) Parse2Command() (cmd *Command, err error) {
	err = json.Unmarshal(m.Data, cmd)
	return
}
