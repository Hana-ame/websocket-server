package main

import (
	"fmt"
	"testing"

	"github.com/Hana-ame/websocket-server/msg"
)

func TestC2S_Regist(t *testing.T) {
	r := &msg.C2S_Regist{
		Name:     "sdfsd",
		Account:  "sdfsdfsddf",
		Password: "fsdfs",
	}
	b, _ := r.JSON()
	c2s := &C2S_Message{
		NetworkMessageID: C2S_Regist,
		Payload:          string(b),
	}
	b, _ = c2s.JSON()
	fmt.Println(string(b))
}

func TestC2S_Login(t *testing.T) {
	r := &msg.C2S_Login{
		Account:  "sdfsdfsddf",
		Password: "fsdfs",
	}
	b, _ := r.JSON()
	c2s := &C2S_Message{
		NetworkMessageID: C2S_Login,
		Payload:          string(b),
	}
	b, _ = c2s.JSON()
	fmt.Println(string(b))
}
