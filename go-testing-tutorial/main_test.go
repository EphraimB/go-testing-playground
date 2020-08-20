package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

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

func TestTableMultiply(t *testing.T) {
	var tests = []struct {
		inputX   int
		inputY   int
		expected int
	}{
		{2, 2, 4},
		{1, 1, 1},
		{1, 2, 2},
		{-1, 3, -3},
		{99999, 1, 99999},
	}

	for _, test := range tests {
		if output := multiply(test.inputX, test.inputY); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.inputX, test.inputY, test.expected, output)
		}
	}
}

func TestHttp(t *testing.T) {
	//
	handler := func(w http.ResponseWriter, r *http.Request) {
		// here we write our expected response, in this case, we return a
		// JSON string which is typical when dealing with REST APIs
		io.WriteString(w, "{ \"status\": \"expected service response\"}")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Error("Test failed. Expected status code to be 200.")
	}

	if resp.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Error("Test failed. Expected content type to equal text/plain; charset=utf-8")
	}

	if string(body) != "{ \"status\": \"expected service response\"}" {
		t.Error("Test failed. Expected status to equal expected service response.")
	}
}

func TestSearchHandlerShouldReturn404IfNoSearchQueryIsPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/?search=Testing", nil)
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	searchHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		t.Error("Wrong status code")
	}
	fmt.Println(string(body))

	if string(body) != "{\"Results\":[\"Cutie\",\"Autism\",\"iPhone 12\"]}" {
		t.Error("Wrong body")
	}
}
