package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// enum, pass
func TestMessageID(t *testing.T) {
	fmt.Println(C2S_Regist)
	fmt.Println(S2C_Regist)
	fmt.Println(C2S_Login)
	fmt.Println(S2C_Login)

	fmt.Println(C2S_Regist == 1)
	fmt.Println(S2C_Regist == 2)
	fmt.Println(C2S_Login == 3)
	fmt.Println(S2C_Login == 4)

}

// marshal, pass
func TestMarshal(t *testing.T) {
	o := &C2S_Message{S2C_Regist, (`123\\<><><>>{"123":1231232}`)}
	fmt.Println(o)
	b, e := json.Marshal(&o)
	fmt.Printf("%s\n", b)
	if e != nil {
		t.Error(e)
	}
	o2 := &C2S_Message{}
	fmt.Println(o2)
	json.Unmarshal(b, o2)
	fmt.Println(o2)
}

// marshal no html escape, pass
func TestMarshalNoEscapeHTML(t *testing.T) {
	// encoder := json.NewEncoder(os.Stdout)
	// encoder.SetEscapeHTML(false)

	o := &C2S_Message{S2C_Regist, (`123\\<><><>>{"123":1231232}`)}
	b, e := json.Marshal(o)
	fmt.Printf("%s\n", b)
	if e != nil {
		t.Error(e)
	}
	o2 := &C2S_Message{}
	fmt.Println(o2)
	json.Unmarshal(b, o2)
	fmt.Println(o2)
}
