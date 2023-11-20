package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// the server messageID from client messageID
func SMsgID(id NetworkMessageID) NetworkMessageID {
	switch id {
	case C2S_Login:
		return S2C_Login
	case C2S_Regist:
		return S2C_Regist
	default:
		return id
	}
}

// send json object to websocket conn
func SendMsg(c *websocket.Conn, typ int, s2c any) {
	b, err := s2c.(*S2C_Message).JSON()
	if err != nil {
		msg := &S2C_Message{
			NetworkMessageID:     -1,
			NetworkMessageResult: 0,
			Payload:              err.Error(),
		}
		b, _ := msg.JSON()
		c.WriteMessage(websocket.TextMessage, b)
	}
	c.WriteMessage(typ, b)
}

// marshal json wich no html escape
func JSON(o any) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(o)
	return trim(buffer.Bytes()), err
}

func trim(b []byte) []byte {
	if b[len(b)-1] == '\n' {
		return b[:len(b)-1]
	}
	return b
}
