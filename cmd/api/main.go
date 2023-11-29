package main

import (
	"log"
	"os"

	"github.com/janczizikow/bookshelf/internal/database"
	"github.com/janczizikow/bookshelf/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	DSN := os.Getenv("DB_DSN")

	log.Printf("starting server")
	log.Printf("connecting to the postgres DB")

	db, err := database.Connect(DSN)
	if err != nil {
		log.Fatal("failed to connect to postgres DB")
	}
	defer db.Close()

	log.Println("connected to the postgres DB")

	server := server.New(db)
	err = server.Run()

	if err != nil {
		log.Fatal("failed to start the server")
	}
}
