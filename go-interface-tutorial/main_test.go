// File: main_test.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockShopDB struct{}

// func (m *MockShopDB) CountCustomers(_ time.Time) (int, error) {
// 	return 1000, nil
// }

// func (m *MockShopDB) CountSales(_ time.Time) (int, error) {
// 	return 333, nil
// }

func (m *MockShopDB) CreateBooks() (bool, error) {
	return true, nil
}

func TestCreateBooks(t *testing.T) {
	// Initialize the mock.
	m := &MockShopDB{}
	// Pass the mock to the calculateSalesRate() function.
	sr, err := createBooks(m)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the return value is as expected, based on the mocked
	// inputs.
	exp := "true"
	if sr != exp {
		t.Fatalf("got %v; expected %v", sr, exp)
	}
}

// func TestCalculateSalesRate(t *testing.T) {
// 	// Initialize the mock.
// 	m := &MockShopDB{}
// 	// Pass the mock to the calculateSalesRate() function.
// 	sr, err := calculateSalesRate(m)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Check that the return value is as expected, based on the mocked
// 	// inputs.
// 	exp := "0.33"
// 	if sr != exp {
// 		t.Fatalf("got %v; expected %v", sr, exp)
// 	}
// }

func TestPostgresQueries(t *testing.T) {
	connStr := "host=localhost port=5400 user=docker password=docker sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	word := "Cutie"

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	p := PostgresRepository{
		sdb: &ShopDB{db},
	}

	createBooks(p.sdb)

	// p.sdb.Query("DELETE FROM books WHERE title=$1", word)

	addBookWithTitle(word)

	query := p.search(word)
	p.sdb.Query("DELETE FROM books WHERE title=$1", word)

	if len(query) != 1 {
		t.Error("Wrong length of strings")
	}

	if strings.Join(query, "") != word {
		t.Error("Wrong title")
	}
}

func addBookWithTitle(title string) {
	connStr := "host=localhost port=5400 user=docker password=docker sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	p := PostgresRepository{
		sdb: &ShopDB{db},
	}

	p.sdb.Query("INSERT INTO books VALUES ($1)", title)
}

func TestHandlerWithNoQuery(t *testing.T) {
	connStr := "host=localhost port=5400 user=docker password=docker sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	api := API{
		repository: PostgresRepository{
			sdb: &ShopDB{db},
		},
	}
	api.searchHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	if resp.Status != "400 Bad Request" {
		t.Error("Supposed to be no query")
	}
}

func TestHandlerWithNoResults(t *testing.T) {
	connStr := "host=localhost port=5400 user=docker password=docker sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	req, err := http.NewRequest("GET", "/?search=testing", nil)
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	api := API{
		repository: PostgresRepository{
			sdb: &ShopDB{db},
		},
	}
	api.searchHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	type Result struct {
		Results []string `json:"results"`
		Count   int      `json:"count"`
	}

	var result Result

	test := json.Unmarshal(body, &result)
	if test != nil {
		t.Error("Unmarshal error")
	}

	if result.Count != 0 {
		t.Error("Expected 0 results. Results not 0.")
	}

	if resp.Status != "200 OK" {
		t.Error("Supposed to be a query")
	}
}

func TestHandlerWithSeveralResults(t *testing.T) {
	connStr := "host=localhost port=5400 user=docker password=docker sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	req, err := http.NewRequest("GET", "/?search=testing", nil)
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	api := API{
		repository: PostgresRepository{
			sdb: &ShopDB{db},
		},
	}
	api.repository.sdb.Query("INSERT INTO books VALUES ($1)", "testing")
	api.repository.sdb.Query("INSERT INTO books VALUES ($1)", "testingcute")
	api.repository.sdb.Query("INSERT INTO books VALUES ($1)", "testingpinch")
	api.searchHandler(w, req)
	api.repository.sdb.Query("DELETE FROM books WHERE title=$1", "testing")
	api.repository.sdb.Query("DELETE FROM books WHERE title=$1", "testingcute")
	api.repository.sdb.Query("DELETE FROM books WHERE title=$1", "testingpinch")

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	type Result struct {
		Results []string `json:"results"`
		Count   int      `json:"count"`
	}

	var result Result

	test := json.Unmarshal(body, &result)
	if test != nil {
		t.Error("Unmarshal error")
	}

	if result.Count != 3 {
		t.Error("Expected 3 results. Results not 3.")
	}

	if strings.Join(result.Results, ",") != "testing,testingcute,testingpinch" {
		t.Error("Expected testing, testingcute, and testingpinch")
	}
}
