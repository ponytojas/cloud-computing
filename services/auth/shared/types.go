package shared

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
