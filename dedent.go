package dedent

import (
	"fmt"
)

func D(s string) string {
	return ""
}

func Dedentf(s string, a ...any) string {
	return D(fmt.Sprintf(s, a...))
}
