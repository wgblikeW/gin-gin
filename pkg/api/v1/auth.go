package v1

import "encoding/json"

type Response struct {
	Allowed bool   `json:"allowed"`
	Denied  bool   `json:"denied,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (rsp *Response) ToString() string {
	data, _ := json.Marshal(rsp)

	return string(data)
}
