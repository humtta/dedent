package dedent

import (
	"fmt"
)

func D(s string) string {
	return ""
}

func Df(s string, a ...any) string {
	return D(fmt.Sprintf(s, a...))
}

func indentOf(s string) string {
	i := 0
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	return s[:i]
}

func isBlank(s string) bool {
	r := s[len(indentOf(s)):]
	return r == "" || r == "\n"
}
