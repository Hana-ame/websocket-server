package msg

import (
	"bytes"
	"encoding/json"
)

type C2S_Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// 返回 non-html escaped JSON
func (t *C2S_Login) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

type S2C_Login struct {
	Name string `json:"name"`
}

// 返回 non-html escaped JSON
func (t *S2C_Login) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
