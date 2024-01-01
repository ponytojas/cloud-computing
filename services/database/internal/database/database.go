package database

import (
	"database/internal/logger"
	"database/shared"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type AuthCheck struct {
	UserId   int
	Username string
	Email    string
}

type Purchase struct {
	PurchaseID   int
	UserID       int
	PurchaseDate time.Time
}

type PurchaseItem struct {
	ItemID       int
	PurchaseID   int
	ProductID    int
	Quantity     int
	PricePerUnit float64
	TotalPrice   float64
}

type Invoice struct {
	InvoiceID   int
	UserID      int
	PurchaseID  int
	TotalAmount float64
	CreatedAt   time.Time
}

type Product struct {
	ProductID   int
	Name        string
	Pricing     float64
	Description string
}

type ProductStock struct {
	ProductStockID int
	ProductID      int
	Quantity       int
}

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()
}

func checkDatabases(db *sql.DB) error {
	_, err := db.Exec("SELECT 1 FROM pg_database WHERE datname='auth'")
	if err != nil {
		log.Error("Error checking database", zap.Error(err))
		return err
	}

	// Check if the table exists
	_, err = db.Exec("SELECT 1 FROM pg_tables WHERE tablename='users'")
	if err != nil {
		return err
	}

	return nil
}

func Init() (*sql.DB, error) {
	if os.Getenv("VSCODE_DEBUG") != "true" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Get database configuration from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := checkDatabases(db); err != nil {
		return nil, err
	}

	log.Debug("Connected to database")

	return db, nil
}

func CreateUser(db *sql.DB, user shared.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	var userID int
	createdAt := time.Now()
	err = db.QueryRow(`INSERT INTO "users" (username, email, created_at) VALUES ($1, $2, $3) RETURNING user_id`,
		user.Username, user.Email, createdAt).Scan(&userID)
	if err != nil {
		return 0, err
	}

	_, err = db.Exec(`INSERT INTO "auth" (user_id, password_hash) VALUES ($1, $2)`,
		userID, hashedPassword)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func LoginUser(db *sql.DB, user shared.User) (AuthCheck, error) {
	var userID int
	var username string
	var email string
	var hashedPassword string
	err := db.QueryRow(`SELECT auth.user_id, auth.password_hash, users.username, users.email FROM users INNER JOIN auth ON auth.user_id = users.user_id WHERE username=$1 OR email=$2`,
		user.Username, user.Email).Scan(&userID, &hashedPassword, &username, &email)
	if err != nil {
		return AuthCheck{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return AuthCheck{}, err
	}

	return AuthCheck{userID, username, email}, nil

}

func CreatePurchase(db *sql.DB, userID int) (int, error) {
	var purchaseID int
	purchaseDate := time.Now()
	err := db.QueryRow(`INSERT INTO "user_purchases" (user_id, purchase_date) VALUES ($1, $2) RETURNING purchase_id`,
		userID, purchaseDate).Scan(&purchaseID)
	if err != nil {
		return 0, err
	}

	return purchaseID, nil
}

func AddPurchaseItem(db *sql.DB, purchaseID int, productID int, quantity int, pricePerUnit float64) (int, error) {
	var itemID int
	totalPrice := float64(quantity) * pricePerUnit
	err := db.QueryRow(`INSERT INTO "purchase_items" (purchase_id, product_id, quantity, price_per_unit, total_price) VALUES ($1, $2, $3, $4, $5) RETURNING item_id`,
		purchaseID, productID, quantity, pricePerUnit, totalPrice).Scan(&itemID)
	if err != nil {
		return 0, err
	}

	return itemID, nil
}

func CreateInvoice(db *sql.DB, userID int, purchaseID int, totalAmount float64) (int, error) {
	var invoiceID int
	createdAt := time.Now()
	err := db.QueryRow(`INSERT INTO "invoices" (user_id, purchase_id, total_amount, created_at) VALUES ($1, $2, $3, $4) RETURNING invoice_id`,
		userID, purchaseID, totalAmount, createdAt).Scan(&invoiceID)
	if err != nil {
		return 0, err
	}

	return invoiceID, nil
}

func CreateProduct(db *sql.DB, product Product) (int, error) {
	var productID int
	createdAt := time.Now()
	err := db.QueryRow(`INSERT INTO "product" (name, pricing, description, created_at) VALUES ($1, $2, $3, $4) RETURNING product_id`,
		product.Name, product.Pricing, product.Description, createdAt).Scan(&productID)
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func GetProduct(db *sql.DB, productID int) (Product, error) {
	var product Product
	err := db.QueryRow(`SELECT * FROM "product" WHERE product_id=$1`, productID).Scan(&product.ProductID, &product.Name, &product.Pricing, &product.Description)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func RemoveProduct(db *sql.DB, productID int) error {
	_, err := db.Exec(`DELETE FROM "product" WHERE product_id=$1`, productID)
	if err != nil {
		return err
	}

	return nil
}

func AddProductStock(db *sql.DB, productID int, quantity int) (int, error) {
	var productStockID int
	createdAt := time.Now()
	err := db.QueryRow(`INSERT INTO "product_stock" (product_id, quantity, created_at) VALUES ($1, $2, $3) RETURNING product_stock_id`,
		productID, quantity, createdAt).Scan(&productStockID)
	if err != nil {
		return 0, err
	}

	return productStockID, nil
}

func UpdateProductStock(db *sql.DB, productID int, newQuantity int) error {
	_, err := db.Exec(`UPDATE "product_stock" SET quantity=$1 WHERE product_id=$2`, newQuantity, productID)
	if err != nil {
		return err
	}

	return nil
}

func UpsertProductStock(db *sql.DB, productID int, quantity int) error {
	var productStockID int
	err := db.QueryRow(`SELECT product_stock_id FROM "product_stock" WHERE product_id=$1`, productID).Scan(&productStockID)
	if err != nil {
		// If the product stock doesn't exist, create it
		if err == sql.ErrNoRows {
			_, err := AddProductStock(db, productID, quantity)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// If the product stock exists, update it
		err := UpdateProductStock(db, productID, quantity)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetProductStock(db *sql.DB, productID int) (int, error) {
	var quantity int
	err := db.QueryRow(`SELECT quantity FROM "product_stock" WHERE product_id=$1`, productID).Scan(&quantity)
	if err != nil {
		return 0, err
	}

	return quantity, nil
}

func GetAllProductStock(db *sql.DB) ([]ProductStock, error) {
	var productStocks []ProductStock
	rows, err := db.Query(`SELECT * FROM "product_stock"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productStock ProductStock
		err := rows.Scan(&productStock.ProductStockID, &productStock.ProductID, &productStock.Quantity)
		if err != nil {
			return nil, err
		}
		productStocks = append(productStocks, productStock)
	}

	return productStocks, nil
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	var products []Product
	rows, err := db.Query(`SELECT * FROM "product"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ProductID, &product.Name, &product.Pricing, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func GetAllProductsWithStock(db *sql.DB) ([]Product, error) {
	var products []Product
	rows, err := db.Query(`SELECT * FROM "product INNER JOIN product_stock ON product.product_id = product_stock.product_id"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ProductID, &product.Name, &product.Pricing, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
