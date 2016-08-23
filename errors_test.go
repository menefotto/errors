package errors

import (
	"testing"
)

func TestErrors(t *testing.T) {
	e := New("custom error")
	t.Log(e)
}
