package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth
var logger *log.Logger

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if creds.Login == "1" && creds.Password == "1" {
		claims := jwt.MapClaims{
			"user_id": 1,
			"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires after 72 hours
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte("your_secret"))

		// Create a response
		response := map[string]string{
			"status": "Login successful",
			"token":  tokenString,
			"expiry": time.Now().Add(time.Hour * 72).String(),
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error creating response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
	}
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println("Change password endpoint hit")

}
