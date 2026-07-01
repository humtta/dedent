package dedent

import (
	"testing"
)

func TestD(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "",
			want: "",
		},
		{
			s:    "  ",
			want: "",
		},
		{
			s:    "\t",
			want: "",
		},
		{
			s:    "lorem",
			want: "lorem",
		},
		{
			s:    "  lorem",
			want: "lorem",
		},
		{
			s:    "\tlorem",
			want: "lorem",
		},
		{
			s:    "\nlorem",
			want: "lorem",
		},
		{
			s:    "  \nlorem",
			want: "lorem",
		},
		{
			s:    "\t\nlorem",
			want: "lorem",
		},
		{
			s:    "\n\nlorem",
			want: "\nlorem",
		},
		{
			s:    "lorem\nipsum",
			want: "lorem\nipsum",
		},
		{
			s:    "lorem\n  ipsum",
			want: "lorem\n  ipsum",
		},
		{
			s:    "lorem\n\tipsum",
			want: "lorem\n\tipsum",
		},
		{
			s:    "  lorem\n    ipsum",
			want: "lorem\n  ipsum",
		},
		{
			s:    "\tlorem\n\t\tipsum",
			want: "lorem\n\tipsum",
		},
		{
			s:    "  lorem\nipsum",
			want: "  lorem\nipsum",
		},
		{
			s:    "\tlorem\nipsum",
			want: "\tlorem\nipsum",
		},
		{
			s:    "    lorem\n      ipsum\n  dolor",
			want: "  lorem\n    ipsum\ndolor",
		},
		{
			s:    "\t\tlorem\n\t\t\tipsum\n\tdolor",
			want: "\tlorem\n\t\tipsum\ndolor",
		},
		{
			s:    "  lorem\n  \n  ipsum",
			want: "lorem\n\nipsum",
		},
		{
			s:    "\tlorem\n\t\n\tipsum",
			want: "lorem\n\nipsum",
		},
		{
			s:    "lorem\n  \n    ",
			want: "lorem\n\n",
		},
		{
			s:    "lorem\n\t\n\t\t",
			want: "lorem\n\n",
		},
		{
			s:    "  \tlorem\n  \tipsum",
			want: "lorem\nipsum",
		},
		{
			s:    "\t  lorem\n\t  ipsum",
			want: "lorem\nipsum",
		},
		{
			s:    "  lorem\n\tipsum",
			want: "  lorem\n\tipsum",
		},
		{
			s:    "  lorem\n\t\t\n  ipsum",
			want: "lorem\n\nipsum",
		},
		{
			s:    "\tlorem\n    \n\tipsum",
			want: "lorem\n\nipsum",
		},
	}

	for _, tt := range tests {
		got := D(tt.s)
		if got != tt.want {
			t.Errorf("D(%q)\ngot: %q\nwant: %q", tt.s, got, tt.want)
		}
	}
}

func TestDf(t *testing.T) {
	tests := []struct {
		s    string
		a    []any
		want string
	}{
		{
			s:    "  %s\n    %s",
			a:    []any{"lorem", "ipsum"},
			want: "lorem\n  ipsum",
		},
	}

	for _, tt := range tests {
		got := Df(tt.s, tt.a...)
		if got != tt.want {
			t.Errorf("Df(%q, %v)\ngot: %q\nwant: %q", tt.s, tt.a, got, tt.want)
		}
	}
}
