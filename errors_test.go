package errors

import (
	"testing"
)

func tester() *errors {
	return New("call indirection")
}
func TestErrors(t *testing.T) {
	e := New("custom error")
	t.Log(e.Error())
	t.Log(tester())
}
