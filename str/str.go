package str

import (
	"strings"
)

type String string

func (s *String) Split(sep string) []string {
	return strings.Split(string(*s), sep)
}
