package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// 在 websocket 中传送的包的类型标志。
type NetworkMessageID int

const (
	C2S_Regist NetworkMessageID = iota //注册
	S2C_Regist
	C2S_Login //登录
	S2C_Login
)

type NetworkMessageResult int

const (
	Unknown            NetworkMessageResult = iota
	OK                                      //成功
	LoginNoAccount                          //用户名不存在
	LoginWrongPassword                      //密码错误
)

// 在 websocket 中传送的包。作为接收。
type C2S_Message struct {

	// 在 websocket 中传送的包的类型标志。
	NetworkMessageID `json:"id"`

	// 以 string 方式编码的 JSON (not base64)
	Payload string `json:"content"`
}

// 返回 non-html escaped JSON
func (t *C2S_Message) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// 在 websocket 中传送的包。作为返回。
type S2C_Message struct {

	// 在 websocket 中传送的包的类型标志。
	NetworkMessageID `json:"id"`

	// 返回的结果类型
	NetworkMessageResult `json:"result"`

	// 以 string 方式编码的 JSON (not base64)
	Payload string `json:"content"`
}

// 返回 non-html escaped JSON
func (t *S2C_Message) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// 状态和 payload
func (s2c *S2C_Message) Load(tpy NetworkMessageResult, e string) *S2C_Message {
	s2c.NetworkMessageResult = tpy
	s2c.Payload = e
	return s2c
}

func (s2c *S2C_Message) SendToConnWithType(c *websocket.Conn, typ int) *S2C_Message {
	SendMsg(c, typ, s2c)
	b, _ := s2c.JSON()
	log.Printf("sent: %s", b)
	return s2c
}

func (s2c *S2C_Message) SetMessageID(id NetworkMessageID) *S2C_Message {
	s2c.NetworkMessageID = id
	return s2c
}
