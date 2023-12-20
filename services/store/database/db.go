package database

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type Stock struct {
	StockId string
	Stock   int
	Product string
	Pricing float64
}

func InitDB(dataSourceName string) (*sql.DB, error) {
	// Try to connect to the database
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("DB connected")

	// Check if the database and table exist
	if err := checkDatabases(db); err != nil {
		return nil, err
	}

	return db, nil
}

func checkDatabases(db *sql.DB) error {
	// Check if the database exists
	_, err := db.Exec("SELECT 1 FROM pg_database WHERE datname='auth'")
	if err != nil {
		return err
	}

	// Check if the table exists
	_, err = db.Exec("SELECT 1 FROM pg_tables WHERE tablename='stocks'")
	if err != nil {
		return err
	}

	return nil
}

func createId(product string) string {
	var builder strings.Builder
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for _, char := range product {
		if !vowels[char] {
			builder.WriteRune(char)
		}
	}
	return strings.ToUpper(builder.String())
}

func AddToStock(db *sql.DB, stock Stock) (bool, error) {
	id := createId(stock.Product)
	product := stock.Product
	amount := stock.Stock
	price := stock.Pricing
	_, err := db.Exec("INSERT INTO stocks (stock_id, product, stock, price) VALUES ($1, $2, $3, $4) ON CONFLICT (stock_id) DO UPDATE SET stock = stocks.stock + $3", id, product, amount, price)
	if err != nil {
		return false, err
	}

	return true, nil
}

func SetPrice(db *sql.DB, price float64, product string) (bool, error) {
	id := createId(product)
	_, err := db.Exec("UPDATE stocks SET pricing = $1 WHERE stock_id = $2", price, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetStock(db *sql.DB, product string) (Stock, error) {
	id := createId(product)
	var stock Stock
	err := db.QueryRow("SELECT * FROM stocks WHERE stock_id = $1", id).Scan(&stock)
	if err != nil {
		return Stock{}, err
	}

	return stock, nil
}

func RemoveStock(db *sql.DB, stock int, product string) (bool, error) {
	id := createId(product)
	_, err := db.Exec("UPDATE stocks SET stock = stock - $1 WHERE stock_id = $2", stock, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RemoveId(db *sql.DB, id string) (bool, error) {
	_, err := db.Exec("DELETE FROM stocks WHERE stock_id = $1", id)
	if err != nil {
		return false, err
	}

	return true, nil
}
