package main

import (
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

func SendMsg(c *websocket.Conn, typ int, s2c any) {
	b, err := json.Marshal(s2c)
	if err != nil {
		c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	}
	c.WriteMessage(typ, b)
}
