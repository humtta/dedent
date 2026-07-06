// Package dedent provides functions for removing shared indentation from
// multiline strings.
package dedent

import (
	"fmt"
	"strings"
)

func D(s string) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 && isBlank(s[:i]) {
		s = s[i+1:]
	}

	var (
		prefix   string
		blank    int
		nonBlank int
	)

	for i := range strings.Lines(s) {
		indent := indentOf(i)
		if isBlank(i) {
			blank += len(indent)
			continue
		}
		if nonBlank == 0 {
			prefix = indent
		} else {
			prefix = commonPrefix(prefix, indent)
		}
		nonBlank++
	}

	if len(prefix) == 0 && blank == 0 {
		return s
	}

	var b strings.Builder
	b.Grow(len(s) - blank - nonBlank*len(prefix))

	for i := range strings.Lines(s) {
		if isBlank(i) {
			i = i[len(indentOf(i)):]
		} else {
			i = i[len(prefix):]
		}
		b.WriteString(i)
	}

	return b.String()
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

func commonPrefix(a, b string) string {
	if len(b) < len(a) {
		a, b = b, a
	}
	for i := range len(a) {
		if a[i] != b[i] {
			return a[:i]
		}
	}
	return a
}
