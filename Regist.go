package main

import (
	"encoding/json"

	"github.com/Hana-ame/websocket-server/db"
	"github.com/Hana-ame/websocket-server/msg"
)

func Regist(payload string) (s2c *S2C_Message) {
	s2c = &S2C_Message{NetworkMessageID: S2C_Regist}

	// receive
	o := &msg.C2S_Regist{}
	if err := json.Unmarshal([]byte(payload), o); err != nil {
		// error
		s2c.Load(Unknown, err.Error())
		return
	}

	// db
	user := &db.User{
		Name:     o.Name,
		Account:  o.Account,
		Password: Sha256(o.Password),
	}
	if err := db.CreateUser(user); err != nil {
		// error
		s2c.Load(Unknown, err.Error())
		return
	}

	// return
	r := &msg.S2C_Regist{}
	// public code
	b, err := r.JSON()
	if err != nil {
		// error
		s2c.Load(Unknown, err.Error())
		return
	}

	s2c.Load(OK, string(b))
	return
}
