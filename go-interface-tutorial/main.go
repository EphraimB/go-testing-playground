// File: main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

type Repository interface {
	search(query string) []string
}

type PostgresRepository struct {
	sdb *ShopDB
}

type FakeDBRepository struct {
	searchQuery []string
}

type API struct {
	repository PostgresRepository
}

func (api *API) searchHandler(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("search")
	fmt.Println("Param1 is: " + param1)
	w.Header().Set("Content-Type", "application/json")

	if param1 == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error - Bad request - 400"))

		return
	}

	type Result struct {
		Results []string `json:"results"`
		Count   int      `json:"count"`
	}

	result := Result{
		Results: api.repository.search(param1),
		Count:   len(api.repository.search(param1)),
	}

	res, _ := json.Marshal(result)
	w.Write(res)

	fmt.Println(string(res))
}

func main() {
	connStr := "host=localhost port=5400 user=docker password=docker sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shopDB := &ShopDB{db}

	m1 := &API{
		repository: PostgresRepository{
			sdb: shopDB,
		},
	}
	//sr, err := calculateSalesRate(shopDB)
	sr, err := createBooks(shopDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", sr)

	http.HandleFunc("/", m1.searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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

func (p PostgresRepository) search(query string) []string {
	rows, err := p.sdb.Query("SELECT * FROM books WHERE title LIKE $1", query+"%")

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
