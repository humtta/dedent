package dedent

import (
	"testing"
)

func TestDedent(t *testing.T) {
	tests := []struct {
		name string
		want string
		s    string
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dedent(tt.s)
			if got != tt.want {
				t.Errorf("Dedent(%q)\ngot: %q\nwant: %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestDedentf(t *testing.T) {
	tests := []struct {
		name string
		want string
		s    string
		a    []any
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dedentf(tt.s, tt.a...)
			if got != tt.want {
				t.Errorf("Dedent(%q, %v)\ngot: %q\nwant: %q", tt.s, tt.a, got, tt.want)
			}
		})
	}
}
