package db_test

import (
	"database/sql"

	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateDb(t *testing.T) {
	db, err := sql.Open("mysql", "root:sasdoP123@tcp(127.0.0.1:3306)/data")
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	
	rows, err := db.Query("SHOW TABLES LIKE 'users'")
	if err != nil {
		t.Fatalf("Error querying database: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Error("Users table does not exist")
	}

	
	var tableName string
	if err := rows.Scan(&tableName); err != nil {
		t.Fatalf("Error scanning row: %v", err)
	}
	if tableName != "users" {
		t.Errorf("Expected table name 'users', got '%s'", tableName)
	}


}
