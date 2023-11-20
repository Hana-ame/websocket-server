package main

import (
	"encoding/json"

	"github.com/Hana-ame/websocket-server/db"
	"github.com/Hana-ame/websocket-server/msg"
)

func Login(payload string) (NetworkMessageResult, string) {
	// s2c = &S2C_Message{NetworkMessageID: S2C_Login}

	o := &msg.C2S_Login{}
	if err := json.Unmarshal([]byte(payload), o); err != nil {
		// error
		// s2c.Load(Unknown, err.Error())
		return Unknown, err.Error()
	}

	// db
	user := &db.User{
		Account: o.Account,
	}
	if err := db.ReadUser(user); err != nil {
		// s2c.Load(LoginNoAccount, err.Error())
		return LoginNoAccount, err.Error()
	}
	if user.Password != Sha256(o.Password) {
		// s2c.Load(LoginWrongPassword, "password not match")
		return LoginWrongPassword, "password not match"
	}

	// return
	r := &msg.S2C_Login{
		Name: user.Name,
	}
	b, err := JSON(r)
	if err != nil {
		// s2c.Load(Unknown, err.Error())
		return Unknown, err.Error()
	}

	//
	// s2c.Load(OK, string(b))
	return OK, string(b)
}
