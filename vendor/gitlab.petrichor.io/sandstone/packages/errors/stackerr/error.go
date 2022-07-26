package stackerr

import (
	"encoding/gob"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type Error struct {
	Message string
	Child   error // this error is holding another error
	Root    error // the error that started it all
	Hide    bool
	stack   *stack
}

type flattenOptions struct {
	sanitize bool
}

type wrapper interface {
	Cause() error
}

func init() {
	gob.Register(Error{})
}

func Newf(message string, args ...any) error {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	err := &Error{
		Message: message,
		stack: &stack{
			frames: make([]uintptr, 5),
			limit:  5,
		},
	}

	runtime.Callers(2, err.stack.frames)

	return err
}

func (e *Error) Error() string {
	out := make([]string, 0, 5)

	e.flatten(&out, nil)

	// reverse
	for i := len(out)/2 - 1; i >= 0; i-- {
		opp := len(out) - 1 - i
		out[i], out[opp] = out[opp], out[i]
	}

	return strings.Join(out, ": ")
}

func (e *Error) flatten(out *[]string, opts *flattenOptions) {
	defer func() {
		if r := recover(); r != nil {
			if e.Root != nil {
				e.Stacktrace().Output()

				panic(fmt.Sprintf("issue while flattening error (%s)", e.Root.Error()))
			} else {
				e.Stacktrace().Output()

				panic(fmt.Sprintf("issue while flattening error (%s)", e.Message))
			}
		}
	}()

	if opts != nil && opts.sanitize {
		if e.Hide {
			return
		}
	}

	if e.Child != nil {
		if se, ok := e.Child.(*Error); ok {
			se.flatten(out, opts)
		} else {
			*out = append(*out, e.Child.Error())
		}
	}

	if e.Message != "" {
		*out = append(*out, e.Message)
	}
}

func Root(in error) (out error) {
	if e, ok := in.(*Error); ok {
		if e.Root != nil {
			return e.Root
		}

		return e
	}

	return in
}

func Cause(in error) (out error) {
	if e, ok := in.(*Error); ok {
		if e.Root != nil {
			if e, ok := e.Root.(wrapper); ok {
				return e.Cause()
			}

			return e.Root
		}

		return e
	}

	if e, ok := in.(wrapper); ok {
		return e.Cause()
	}

	return in
}

func Sanitize(in error) (out error) {
	if e, ok := in.(*Error); ok {
		out := make([]string, 0, 5)

		opts := &flattenOptions{
			sanitize: true,
		}

		e.flatten(&out, opts)

		// reverse
		for i := len(out)/2 - 1; i >= 0; i-- {
			opp := len(out) - 1 - i
			out[i], out[opp] = out[opp], out[i]
		}

		return errors.New(strings.Join(out, ": "))
	}

	return in
}
