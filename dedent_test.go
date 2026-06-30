package dedent

import (
	"testing"
)

func TestD(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := D(tt.s)
			if got != tt.want {
				t.Errorf("D(%q)\ngot: %q\nwant: %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestDf(t *testing.T) {
	tests := []struct {
		name string
		s    string
		a    []any
		want string
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Df(tt.s, tt.a...)
			if got != tt.want {
				t.Errorf("Df(%q, %v)\ngot: %q\nwant: %q", tt.s, tt.a, got, tt.want)
			}
		})
	}
}
