package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

func additionalHandler(search string) string {
	var query string
	handler := func(w http.ResponseWriter, r *http.Request) string {
		query = r.URL.Query().Get("result")
		return query
	}
	req, err := http.NewRequest("GET", "/?result="+search, nil)
	if err != nil {
		return "Error"
	}
	w := httptest.NewRecorder()
	handler(w, req)

	//resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	return query
}

func main() {
	fmt.Println(additionalHandler("Testing"))
	fmt.Println("Hello World")
}
