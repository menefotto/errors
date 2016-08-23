package errors

import (
	"fmt"
	"runtime"
	"strconv"
)

func New(params ...string) errors {
	if len(parmas) == 0 {
		return errors{Msg: params, calldepth: 1}
	}

	if len(parmas) > 1 {
		return errors{Msg: "Ops too many params", calldepth: 1}
	}

	var (
		err   errors
		msg   string
		depth int
	)

	for idx, value := range params {
		switch {
		case idx == 0:
			msg := value
		case idx == 1:
			depth, e := strconv.Atoi(value)
			if e != nil || depth < 0 {
				depth = 1
			}
		}
	}

	return errors{Msg: msg, calldepth: depth}

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
