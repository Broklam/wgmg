package main

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = GenerateRandomSecretKey(32)

func main() {
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/change_password", changePasswordHandler)
	http.HandleFunc("/api/test", authenticationMiddleware(testHandler))

	// Start the server
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("Error listening on port :8080:", err)
		}
	}()

	log.Println("Server started on :8080")
	select {}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "test" && password == "test" {
		token, expiration := generateJWT(username)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     "successful",
			"token":      token,
			"expires_in": expiration.Format(time.RFC3339),
		})
		log.Printf("Login successful for user: %s", username)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		log.Printf("Login failed for user: %s", username)
	}
}

func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	//pass logic
	// Respond with success message
	json.NewEncoder(w).Encode(map[string]string{"message": "Password changed successfully"})
	log.Println("Password changed successfully")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "I AM STILL NOT FINISHED"})
	log.Println("Test endpoint accessed")
}

func generateJWT(username string) (string, time.Time) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString, time.Now().Add(time.Hour * 24)
}

func authenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Future auth logic (e.g., verify token from request header)
		next(w, r)
	}
}

func GenerateRandomSecretKey(length int) []byte {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Error generating secret key:", err)
	}
	return randomBytes
}
