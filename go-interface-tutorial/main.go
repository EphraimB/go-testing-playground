package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// // Declare a Book type which satisfies the fmt.Stringer interface.
// type Book struct {
// 	Title  string
// 	Author string
// }

// func (b Book) String() string {
// 	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
// }

// // Declare a Count type which satisfies the fmt.Stringer interface.
// type Count int

// func (c Count) String() string {
// 	return strconv.Itoa(int(c))
// }

// // Declare a WriteLog() function which takes any object that satisfies
// // the fmt.Stringer interface as a parameter.
// func WriteLog(s fmt.Stringer) {
// 	log.Println(s.String())
// }

// // Create a Customer type
// type Customer struct {
// 	Name string
// 	Age  int
// }

// // Implement a WriteJSON method that takes an io.Writer as the parameter.
// // It marshals the customer struct to JSON, and if the marshal worked
// // successfully, then calls the relevant io.Writer's Write() method.
// func (c *Customer) WriteJSON(w io.Writer) error {
// 	js, err := json.Marshal(c)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = w.Write(js)
// 	return err
// }

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

func main() {
	// 	// Initialize a Count object and pass it to WriteLog().
	// 	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	// 	WriteLog(book)

	// 	// Initialize a Count object and pass it to WriteLog().
	// 	count := Count(3)
	// 	WriteLog(count)

	// 	// Initialize a customer struct.
	// 	c := &Customer{Name: "Alice", Age: 21}

	// 	// We can then call the WriteJSON method using a buffer...
	// 	var buf bytes.Buffer
	// 	err := c.WriteJSON(&buf)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// Or using a file.
	// 	f, err := os.Create("/tmp/customer")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer f.Close()

	// 	err = c.WriteJSON(f)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shopDB := &ShopDB{db}
	sr, err := calculateSalesRate(shopDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sr)
}

func calculateSalesRate(sdb *ShopDB) (string, error) {
	since := time.Now().Add(-24 * time.Hour)

	sales, err := sdb.CountSales(since)
	if err != nil {
		return "", err
	}

	customers, err := sdb.CountCustomers(since)
	if err != nil {
		return "", err
	}

	rate := float64(sales) / float64(customers)
	return fmt.Sprintf("%.2f", rate), nil
}
