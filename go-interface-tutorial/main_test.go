// File: main_test.go
package main

import "testing"

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

type TestRepository struct {
}

func (t *TestRepository) search(query string) []string {
	return []string{"Star Wars", "Harry Potter"}
}

func TestPostgresQueries(t *testing.T) {
	p := PostgresRepository{
		repository: &TestRepository{},
	}
	p.search("Testing")
}
