package types

import "encoding/json"

type RequestWrapper struct {
	Auth    *Auth           `json:"auth"`
	Request json.RawMessage `json:"request"`
}

func (rw RequestWrapper) Unmarshal(out interface{}) error {
	return json.Unmarshal(rw.Request, out)
}
