package helloworld

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"world", "Hello world"},
		{"世界", "Hello 世界"},
		{"2 words", "Hello 2 words"},
	}
	for _, c := range cases {
		got := SayHello(c.in)
		if got != c.want {
			t.Errorf("Helloworld(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
