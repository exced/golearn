package main

import (
	"testing"
)

func TestWrapper(t *testing.T) {
	want := []int{1, 2, 3}
	increment := wrapper()
	for i := range want {
		got := increment()
		if got != want[i] {
			t.Errorf("wrapper() == %q, want %q", got, want[i])
		}
	}
}
