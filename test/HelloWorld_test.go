package example_test

import (
	"testing"

	"example/pkg/example"
)

/**
* NOTE: Welcome to Golang Test Unit!
*
* This directory contains all of your test units file.
* Every files suffixed as `_test.go` is considered a test unit files.
* And all Function contains args (t *testing.T) will be tested.
*
* Feel free to restructure your code as you see fit.
* Try running this test unit with `make test`.
**/
func TestAdd(t *testing.T) {
	// Test cases
	tests := []struct {
		a, b   int
		result int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{-1, -1, -2},
		{-1, 1, 0},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		got := example.Add(tt.a, tt.b)
		if got != tt.result {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.result)
		}
	}
}
