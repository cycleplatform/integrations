package stackerr

import "runtime"

func Hide(err error) (error) {
	e := &Error{
		Child: err,
		Hide:  true,
	}

	if se, ok := err.(*Error); ok {
		if se.Root != nil {
			e.Root = se.Root
		} else {
			e.Root = err
		}

		if se.stack != nil {
			e.stack = se.stack
		}
	} else {
		e.Root = err
	}

	if e.stack == nil {
		e.stack = &stack{
			frames: make([]uintptr, 5),
			limit:  5,
		}

		runtime.Callers(2, e.stack.frames)
	}

	return e
}
