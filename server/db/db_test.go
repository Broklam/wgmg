package db_test

import (
	"database/sql"

	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateDb(t *testing.T) {
	// Replace with your actual MySQL connection details
	db, err := sql.Open("mysql", "root:sasdoP123@tcp(127.0.0.1:3306)/data")
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Check if the users table exists
	rows, err := db.Query("SHOW TABLES LIKE 'users'")
	if err != nil {
		t.Fatalf("Error querying database: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Error("Users table does not exist")
	}

	// Verify the table structure (columns)
	// You can add more checks here based on your expected schema

	// Example: Check if the 'username' column exists
	var tableName string
	if err := rows.Scan(&tableName); err != nil {
		t.Fatalf("Error scanning row: %v", err)
	}
	if tableName != "users" {
		t.Errorf("Expected table name 'users', got '%s'", tableName)
	}

	// You can add more checks for other columns here

	// Feel free to adapt this example to your specific use case!
}
