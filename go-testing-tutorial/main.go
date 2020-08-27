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

type Repository interface {
	search(query string) []string
}

type API struct {
	repository Repository
}

func (api *API) searchHandler(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("search")
	fmt.Println("Param1 is: " + param1)
	if param1 == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)

		// use api.repository here
		type SearchResults struct {
			Results []string `json:"Results"`
		}

		searchResults := SearchResults{
			Results: api.repository.search(param1),
		}
		// ----------------------------------

		var jsonData []byte
		jsonData, err := json.Marshal(searchResults)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(jsonData))

		io.WriteString(w, string(jsonData))
	}
}

func main() {
	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
