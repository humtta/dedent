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
