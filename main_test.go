package main

import "testing"

func TestAdd(t *testing.T) {
	if add(2, 4) != 6 {
		t.Error("Expected 2 + 4 to equal 6")
	}
}

func TestTableAdd(t *testing.T) {
	var tests = []struct {
		inputX   int
		inputY   int
		expected int
	}{
		{2, 2, 4},
		{-1, 2, 1},
		{0, 2, 2},
		{-5, 2, -3},
		{99999, 2, 100001},
	}

	for _, test := range tests {
		if output := add(test.inputX, test.inputY); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.inputX, test.inputY, test.expected, output)
		}
	}
}

func TestSubtract(t *testing.T) {
	if subtract(7, 5) != 2 {
		t.Error("Expected 7 - 5 to equal 2")
	}
}

func TestTableSubtract(t *testing.T) {
	var tests = []struct {
		inputX   int
		inputY   int
		expected int
	}{
		{6, 2, 4},
		{3, 2, 1},
		{4, 2, 2},
		{-1, 2, -3},
		{99999, 2, 99997},
	}

	for _, test := range tests {
		if output := subtract(test.inputX, test.inputY); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.inputX, test.inputY, test.expected, output)
		}
	}
}

func TestMultiply(t *testing.T) {
	if multiply(2, 5) != 10 {
		t.Error("Expected 2 * 5 to equal 10")
	}
}
