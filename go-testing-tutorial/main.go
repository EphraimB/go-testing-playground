package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	param1 := r.URL.Query().Get("search")

	if param1 == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	type SearchResults struct {
		Results []string `json:"Results"`
	}

	searchResults := SearchResults{
		Results: []string{"Cutie", "Autism", "iPhone 12"},
	}

	var jsonData []byte
	jsonData, err := json.Marshal(searchResults)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	io.WriteString(w, string(jsonData))
}

func main() {
	fmt.Println("Hello World")
}
