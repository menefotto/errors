package errors

import (
	"fmt"
	"runtime"
)

func New(err string) errors {
	return errors{Msg: err, calldepth: 1}

}

func New(err string, depth int) errors {
	return errros{Msg: err, calldepth: depth}
}

type errors struct {
	Msg       string
	calldepth int
}

func (e *errors) Error() string {
	var (
		ok   bool
		line int
		file string
	)

	_, file, line, ok = runtime.Caller(e.calldepth)
	if !ok {
		file = "Ops no file name"
		line = -1
	}

	if len(e.Msg) > 79 {
		msg := []byte(e.Msg)
		return fmt.Sprintf("%s\n%s: %s: %v\n",
			msg[:79], msg[79:], file, line)

	}
	return fmt.Sprintf("%s: %s: %v\n", e.Msg, file, line)

}
