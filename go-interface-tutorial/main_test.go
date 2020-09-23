// File: main_test.go
package main

import (
	"database/sql"
	"log"
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

	p.sdb.Query("DELETE FROM books WHERE title='" + word + "'")

	addBookWithTitle(word)

	query := p.search(word)

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

	p.sdb.Query("INSERT INTO books VALUES ('" + title + "')")
}
