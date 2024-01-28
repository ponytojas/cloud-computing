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

type TokenResponse struct {
	Status string
	Token  string
	User   AuthCheck
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

type ProductsResponse struct {
	Status string
	Data   []Product
}

type ProductResponse struct {
	Status  string
	Product Product
}

type ProductCreationResponse struct {
	Id int
}

type ProductStock struct {
	ProductStockID int
	ProductID      int
	Quantity       int
}

type ProductStockResponse struct {
	Status string
	Data   []ProductStock
}
