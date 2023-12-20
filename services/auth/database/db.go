package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

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
	_, err = db.Exec("SELECT 1 FROM pg_tables WHERE tablename='users'")
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(db *sql.DB, user User) (int, error) {
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

func LoginUser(db *sql.DB, user User) (AuthCheck, error) {
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
