package msg

import (
	"bytes"
	"encoding/json"
)

type C2S_Regist struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type S2C_Regist struct {
}

// 返回 non-html escaped JSON
func (t *S2C_Regist) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
