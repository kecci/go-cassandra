package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kecci/go-cassandra/model"
	"github.com/kecci/go-cassandra/repository"
)

// AddUser add
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	if err := repository.AddUser(&user); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	println("User added", user.Email)
	json.NewEncoder(w).Encode("User added")
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if email == "" {
		json.NewEncoder(w).Encode("email is required")
		return
	}

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		log.Println("error getting user", email, err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
