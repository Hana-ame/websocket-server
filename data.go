package main

import (
	"bytes"
	"encoding/json"
)

// 在 websocket 中传送的包的类型标志。
type MessageID int

const (
	Undefined MessageID = iota
	C2S_Regist
	S2C_Regist
	C2S_Login
	S2C_Login
)

// 在 websocket 中传送的包。
type Message struct {

	// 在 websocket 中传送的包的类型标志。
	MessageID

	// 以 string 方式编码的 JSON (not base64)
	Payload string
}

// 返回 non-html escaped JSON
func (t *Message) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
