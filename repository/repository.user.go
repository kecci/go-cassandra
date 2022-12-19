package repository

import (
	"encoding/json"
	"log"

	"github.com/kecci/go-cassandra/database"
	"github.com/kecci/go-cassandra/model"
)

func AddUser(user *model.User) error {
	// query
	query := `INSERT INTO users (email, first_name, last_name) VALUES (?, ?, ?)`

	// execute query
	if err := database.ExecuteQuery(query, user.Email, user.FirstName, user.LastName); err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (user *model.User, err error) {
	// query
	query := `SELECT * FROM users WHERE email = ?`

	// response map string interface
	res, err := database.GetQuery(query, email)
	if err != nil {
		log.Println("Error database.GetQuery", err.Error())
		return nil, err
	}

	// convert map to json
	resByte, err := json.Marshal(res)
	if err != nil {
		log.Println("Error json.Marshal", err.Error())
		return nil, err
	}

	// convert json to struct
	if err = json.Unmarshal(resByte, &user); err != nil {
		log.Println("Error json.Marshal", err.Error())
		return nil, err
	}

	return user, nil
}
