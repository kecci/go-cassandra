package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kecci/go-cassandra/handler"
)

func NewRouter() http.Handler {

	// Router
	r := chi.NewRouter()

	// User
	r.Get("/user/{email}", handler.GetUserByEmail)
	r.Post("/user", handler.AddUser)

	log.Println("running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
	return r
}
