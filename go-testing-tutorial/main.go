package main

import (
	"fmt"
	"io"
	"net/http"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func add(x int, y int) (result int) {
	result = x + y

	return result
}

func subtract(x int, y int) (result int) {
	result = x - y

	return result
}

func multiply(x int, y int) (result int) {
	result = x * y

	return result
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	fmt.Println(query)
	// JSON string which is typical when dealing with REST APIs
	io.WriteString(w, "{ \"status\": \"expected service response\"}")

	param1 := r.URL.Query().Get("search")
	if param1 != "" {
		w.WriteHeader(400)
	} else {
		w.WriteHeader(200)
	}
}

func main() {
	fmt.Println("Hello World")
}
