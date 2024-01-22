package shared

type User struct {
	Username string
	Email    string
	Password string
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

type ProductStockRequest struct {
	Quantity int
}
