package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// User represents a user record in the database
type User struct {
	ID           int
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	ModifiedAt   time.Time
}

func CreateDb() {
	// Replace with your actual MySQL connection details
	db, err := sql.Open("mysql", "root:sasdoP123@tcp(127.0.0.1:3306)/data")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table
	_, err = db.Exec(`

		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			passwordhash VARCHAR(258) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Example CRUD operations:
	// - Create a new user
	// - Read user data
	// - Update user data
	// - Delete a user

	// Implement your CRUD functions here based on your application needs.
	// You can use db.Exec, db.Query, and other methods to interact with the database.

	fmt.Println("Users table created 1successfully!")
}
