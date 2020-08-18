package main

import "testing"

func TestAdd(t *testing.T) {
	if add(2, 4) != 6 {
		t.Error("Expected 2 + 4 to equal 6")
	}
}

func TestSubtract(t *testing.T) {
	if subtract(7, 5) != 2 {
		t.Error("Expected 7 - 5 to equal 2")
	}
}

func TestMultiply(t *testing.T) {
	if multiply(2, 5) != 10 {
		t.Error("Expected 2 * 5 to equal 10")
	}
}
