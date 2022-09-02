package js

import (
	"github.com/hamidreza01/js.go/slice"
	"github.com/hamidreza01/js.go/str"
)

type String str.String

type Slice[T any] slice.Slice[T]
