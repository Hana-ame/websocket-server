package main

import (
	"fmt"
	"testing"

	"github.com/Hana-ame/websocket-server/msg"
)

func TestRegist(t *testing.T) {
	r := &msg.C2S_Regist{
		Name:     "sdfsd",
		Account:  "sdfsdfsaddf",
		Password: "fsdfs",
	}
	b, _ := r.JSON()
	// c2s := &C2S_Message{
	// 	NetworkMessageID: C2S_Regist,
	// 	Payload:          string(b),
	// }
	// b, _ = c2s.JSON()
	s2c := &S2C_Message{}
	s2c.Load(Regist(string(b)))
	b, _ = JSON(s2c)
	fmt.Printf("%s", b)
}
