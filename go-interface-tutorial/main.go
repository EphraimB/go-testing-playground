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
		_, err := sdb.Query("CREATE TABLE books (title VARCHAR(50) PRIMARY KEY)")

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
	sdb *ShopDB
}

type FakeDBRepository struct {
	searchQuery []string
}

func (p PostgresRepository) search(query string) []string {
	rows, err := p.sdb.Query("SELECT * FROM books WHERE title=$1", query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", names)

	return names
}

func (f FakeDBRepository) search() {

}
