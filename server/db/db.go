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
	Role         string
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
			passwordhash VARCHAR(900) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	
	CREATE TABLE IF NOT EXISTS roles (
		user_id INT,
		role VARCHAR(50) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`)

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		regularity VARCHAR(50),
		difficulty INT CHECK (difficulty BETWEEN 1 AND 5),
		importance INT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)
	`)

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS messages (
		message_id INT AUTO_INCREMENT PRIMARY KEY,
		sender_id INT NOT NULL,
		content TEXT NOT NULL,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (sender_id) REFERENCES users(id)
	)
	`)

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS todo (
		id INT AUTO_INCREMENT PRIMARY KEY,
		Title VARCHAR(255) NOT NULL,
		Description TEXT,
		Due_Date DATE,
		Author VARCHAR(255),
		Done BOOLEAN,
		FOREIGN KEY (Author) REFERENCES users(username),
		-- Add any other necessary columns
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)
	`)

	if err != nil {
		log.Fatal(err)
	}

	// future
	fmt.Println("Migration Done")
}
