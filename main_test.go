package main

import (
	"os"
	"testing"
)

// Testing get arguments
func TestGetArguments(t *testing.T) {
	args := os.Args
	if len(args) != 3 {
		t.Error("Expected two arguments")
	} else {
		time := args[1]
		day := os.Args[2]
		if len(time) != 5 {
			t.Error("Expected time formatt: HH:MM")
		}
		if len(day) < 0 {
			t.Error("Expected day greater than ZERO")
		}
	}
}
