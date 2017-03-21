package errors

import (
	"fmt"
	"testing"
)

func tester() error {
	return New("call indirection")
}

func TestErrors(t *testing.T) {
	e := New("custom error")
	t.Log(e.Error())
	t.Log(tester())

	wrapped := Wrap(fmt.Errorf("wrapped error\n"))
	t.Log(wrapped())
}

func TestMsgToLong(t *testing.T) {
	New("thi message is suppose to be very long really really long and it isn't really that long weel let's see")
}

func TestTooManyParmas(t *testing.T) {
	New("thodisa", "fdsajkl", "flkasdll")
}

func TestCustomCallDepth(t *testing.T) {
	New("ciao carlo", "-1")
}
