package main

import (
	"fmt"
	"testing"

	"github.com/Hana-ame/websocket-server/msg"
)

func TestLogin(t *testing.T) {
	r := &msg.C2S_Login{
		Account:  "123",
		Password: "123",
	}
	b, _ := JSON(r)
	s2c := &S2C_Message{}
	s2c.Load(Login(string(b)))
	b, _ = JSON(s2c)
	fmt.Printf("%s", b)
}
