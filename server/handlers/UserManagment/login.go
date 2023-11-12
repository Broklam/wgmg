package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Broklam/wg/hasher"
	"github.com/dgrijalva/jwt-go"

	_ "github.com/go-sql-driver/mysql"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

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

	var passwordHash string
	log.Println(req.Username)
	err = db.QueryRow("SELECT passwordhash FROM users WHERE username = ?", req.Username).Scan(&passwordHash)
	log.Println(passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Username or password is incorrect"})
		} else {
			log.Printf("Database error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Internal server error"})
		}
		return
	}
	match, _ := hasher.ComparePasswordAndHash(req.Password, passwordHash)
	log.Println(match)

	if err != nil || match != true {
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(LoginResponse{Token: tokenString})
}
