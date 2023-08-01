package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mohamedafify/backend/api"
	"github.com/mohamedafify/backend/storage"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the enviroment variables")
	}

	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewServer("localhost:"+port, store)
	server.Start()
}
