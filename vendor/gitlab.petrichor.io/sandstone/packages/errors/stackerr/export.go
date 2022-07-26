package stackerr

import (
	"encoding/gob"
	"encoding/json"
	"reflect"
	"strings"
)

type Exported struct {
	Message         string `bson:"message" json:"message"`
	MessageInternal string `bson:"message_internal" json:"message_internal,omitempty"`
	Stack           Frames `bson:"stack" json:"stack,omitempty"`
}

var SanitizeJSON bool

func (e *Exported) MarshalJSON() ([]byte, error) {
	ne := *e

	if SanitizeJSON {
		ne.MessageInternal = ""
		ne.Stack = nil
	}

	return json.Marshal(ne)
}

func init() {
	gob.Register(Exported{})
}

func Exportable(err error) bool {
	switch err.(type) {
	case *Error:
		return true
	default:
		if strings.Contains(reflect.TypeOf(err).String(), "errors.errorString") {
			return true
		}
		return false
	}
}

func Export(err error) *Exported {
	if err == nil {
		panic("Error is nil")
	}

	switch e := err.(type) {
	case Exported:
		return &e
	case *Exported:
		return e
	case *Error:
		ex := &Exported{
			Message:         Sanitize(err).Error(),
			MessageInternal: err.Error(),
		}

		if e.stack != nil {
			ex.Stack = e.Stacktrace()
		}

		return ex
	default:
		return &Exported{
			Message:         e.Error(),
			MessageInternal: e.Error(),
		}
	}
}

func (e Exported) Error() string {
	return e.Message
}
