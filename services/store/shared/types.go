package shared

type Product struct {
	ProductID   int
	Name        string
	Pricing     float64
	Description string
}

type ProductCreationResponse struct {
	Id int
}

type ProductStock struct {
	ProductStockID int
	ProductID      int
	Quantity       int
}

type ProductStockRequest struct {
	Quantity int
}
