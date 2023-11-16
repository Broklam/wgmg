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
	"github.com/go-chi/cors"
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

	tokenAuth = jwtauth.New("HS256", []byte("jghkvilgu25rj"), nil)

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

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Use this to allow specific origin hosts
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))

	r.Group(func(r chi.Router) {

		r.Post("/api/login", handlers.LoginHandler)
		r.Post("/api/signup", handlers.RegisterNewUser)
	})

	// private routes
	r.Route("/api", func(r chi.Router) {
		// jwt verification
		r.Use(jwtauth.Verifier(tokenAuth))

		// authenticator
		r.Use(jwtauth.Authenticator)

		r.Post("/change_password", handlers.LoginHandler)
		/* what i want to add
		r.Post("/change_role", handlers.ChangeRoleHandler)
		r.Get("/get_todos", handlers.GetTodosHandler)
		r.Post("/add_todo", handlers.AddTodoHandler)
		r.Delete("/delete_todo", handlers.DeleteTodoHandler)
		r.Get("/get_messages", handlers.GetMessagesHandler)
		r.Post("/send_message", handlers.SendMessageHandler)
		r.Get("/get_all_tasks", handlers.GetAllTasksHandler)
		*/
	})

	logger.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", r)
}
