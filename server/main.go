package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/Broklam/wg/db"
	"github.com/Broklam/wg/handlers/UserManagment"

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

	db.CreateDb()
	r := chi.NewRouter()

	//  logging middleware
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))

	// public routes
	// TODO FIX ROUTING ITS A MESS
	r.Group(func(r chi.Router) {
		r.Get("/public", func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Public endpoint hit")
			w.Write([]byte("Public endpoint, no authentication required"))
		})
		r.Post("/login", handlers.LoginHandler)
		r.Post("/signup", handlers.RegisterNewUser)
		r.Post("/deleteuser", handlers.DeleteUser) //TODO change to protected later!!
	})

	// private routes
	r.Group(func(r chi.Router) {
		// jwt verif
		r.Use(jwtauth.Verifier(tokenAuth))

		// authenticator
		r.Use(jwtauth.Authenticator)

		r.Post("/change-password", handlers.LoginHandler)
	})

	logger.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", r)
}
