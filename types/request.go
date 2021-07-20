package types

type RequestWrapper struct {
	Auth    *Auth       `json:"auth"`
	Request interface{} `json:"request"`
}
