package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Broklam/wg/hasher"
	_ "github.com/go-sql-driver/mysql"
)

type RegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "root:sasdoP123@tcp(127.0.0.1:3306)/data")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	username := req.Username
	password := req.Password
	passwordHash := hasher.GenerateFromPassword(password)

	stmt, err := db.Prepare("INSERT INTO users (username, PasswordHash) VALUES (?, ?)")
	if err != nil {
		log.Println("Error preparing SQL statement:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, passwordHash)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User registered successfully!")
}
