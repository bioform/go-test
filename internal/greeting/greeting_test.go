package greeting

import (
	"errors"
	"testing"
)

func TestGreet(t *testing.T) {
	cases := []struct {
		in, want string
		err      error
	}{
		{"mmm", "Hello, mmm", nil},
		{"", "", ErrEmptyName},
	}

	for _, c := range cases {
		got, err := Greet(c.in)
		if got != c.want || !errors.Is(err, c.err) {
			t.Errorf("Greet(%q) == %q, %q, want %q, %q", c.in, got, err, c.want, c.err)
		}
	}
}
