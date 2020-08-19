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
	io.WriteString(w, "{ \"status\": \"something\"}")
}

func main() {
	fmt.Println("Hello World")
}
