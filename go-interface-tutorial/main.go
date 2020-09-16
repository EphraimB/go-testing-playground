// File: main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Create our own custom ShopModel interface. Notice that it is perfectly
// fine for an interface to describe multiple methods, and that it should
// describe input parameter types as well as return value types.
type ShopModel interface {
	//CountCustomers(time.Time) (int, error)
	//CountSales(time.Time) (int, error)
	CreateBooks() (bool, error)
}

// The ShopDB type satisfies our new custom ShopModel interface, because it
// has the two necessary methods -- CountCustomers() and CountSales().
type ShopDB struct {
	*sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func (sdb *ShopDB) CreateBooks() (bool, error) {
	tableCheck, err := sdb.Query("SELECT * FROM books;")

	if tableCheck == nil {
		_, err := sdb.Query("CREATE TABLE books (title VARCHAR(50) PRIMARY KEY);")

		return false, err
	} else {
		return true, err
	}
}

func main() {
	connStr := "host=localhost port=5400 user=docker password=docker sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shopDB := &ShopDB{db}
	//sr, err := calculateSalesRate(shopDB)
	sr, err := createBooks(shopDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", sr)

	//fmt.Println(PostgresRepository.searchTable)
}

func createBooks(sm ShopModel) (string, error) {
	books, err := sm.CreateBooks()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%t", books), nil
}

// Swap this to use the ShopModel interface type as the parameter, instead of the
// concrete *ShopDB type.
// func calculateSalesRate(sm ShopModel) (string, error) {
// 	since := time.Now().Add(-24 * time.Hour)

// 	sales, err := sm.CountSales(since)
// 	if err != nil {
// 		return "", err
// 	}

// 	customers, err := sm.CountCustomers(since)
// 	if err != nil {
// 		return "", err
// 	}

// 	rate := float64(sales) / float64(customers)
// 	return fmt.Sprintf("%.2f", rate), nil
// }

type Repository interface {
	search(query string) []string
}

type PostgresRepository struct {
	repository  Repository
	searchTable []string
	sdb         *ShopDB
}

type FakeDBRepository struct {
	searchQuery []string
}

func (p PostgresRepository) search(query string) []string {
	rows, err := p.sdb.Query("SELECT * FROM books WHERE title='" + query + "'")
	if err != nil {
		fmt.Println("Failed to run query", err)
		return []string{}
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return []string{}
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return []string{}
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)
	}

	return result
}

func (f FakeDBRepository) search() {

}
