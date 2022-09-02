package test

import (
	"github.com/hamidreza01/js.go/slice"
	"github.com/hamidreza01/js.go/str"
	"testing"
)

func TestSplit(t *testing.T) {
	var s str.String = "Hello World"
	var sl slice.Slice[string] = s.Split(" ")
	if len(sl) != 2 {
		t.Errorf("`str.String.Split` method is not working properly, out : %v", sl)
	}
}
