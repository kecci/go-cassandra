package main

import (
	"github.com/kecci/go-cassandra/database"
	"github.com/kecci/go-cassandra/router"
)

func init() {
	database.InitDatabase()
}

func main() {
	router.NewRouter()
}
