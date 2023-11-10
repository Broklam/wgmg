package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth
var logger *log.Logger

type begoneEscapeSign struct {
	w io.Writer
}

func (ef *begoneEscapeSign) Write(p []byte) (n int, err error) {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	cleanedP := re.ReplaceAll(p, nil)
	return ef.w.Write(cleanedP)
}

func init() {
	//TODO fix not so secret key currently
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	//logger init
	file, err := os.OpenFile("./logs/full_log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	filter := &begoneEscapeSign{w: file}
	multi := io.MultiWriter(filter, os.Stdout)
	logger = log.New(multi, "", log.LstdFlags)
}

func main() {
	r := chi.NewRouter()

	//  logging middleware
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))

	// public routes
	r.Group(func(r chi.Router) {
		r.Get("/public", func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Public endpoint hit")
			w.Write([]byte("Public endpoint, no authentication required"))
		})
		r.Post("/login", LoginHandler)
	})

	// private routes
	r.Group(func(r chi.Router) {
		// jwt verif
		r.Use(jwtauth.Verifier(tokenAuth))

		// authenticator
		r.Use(jwtauth.Authenticator)

		r.Post("/change-password", ChangePasswordHandler)
	})

	logger.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", r)
}

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
