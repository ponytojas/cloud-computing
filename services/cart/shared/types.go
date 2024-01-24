package shared

import "github.com/golang-jwt/jwt"

type User struct {
	Username string
	Email    string
	Password string
}

type AuthCheck struct {
	UserId   int
	Username string
	Email    string
}

type Token struct {
	Token string
}

type TokenLogout struct {
	Username string
}

type Response struct {
	Status string
}

type TokenCheckResponse struct {
	Valid  bool
	claims jwt.MapClaims
}

type Product struct {
	ProductID   int
	Name        string
	Pricing     float64
	Description string
}

type ProductResponse struct {
	Status string
	Data   []Product
}

type ProductCreationResponse struct {
	Id int
}

type ProductStock struct {
	ProductStockID int
	ProductID      int
	Quantity       int
}

type ProductStockSingleGetResponse struct {
	Status string
	Stock  ProductStock
}

type ProductStockResponse struct {
	Status string
	Data   []ProductStock
}

type AddToCartRequest struct {
	UserId    int
	ProductId int
	Quantity  int
}
