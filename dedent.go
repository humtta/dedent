package dedent

import (
	"fmt"
)

func Dedent(s string) string {
	return ""
}

func Dedentf(s string, a ...any) string {
	return Dedent(fmt.Sprintf(s, a...))
}
