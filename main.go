package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedafify/backend/db"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	InitUsers(db, router)

	log.Println("server running on Port:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
