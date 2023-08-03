package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedafify/backend/db"
	"github.com/mohamedafify/backend/handlers"
)

func main() {
	dbConnStr := "postgres://postgres@localhost/backend?sslmode=disable"
	db, err := db.NewDB(dbConnStr)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	userHandler := handlers.NewUserHandler(db)

	router.HandleFunc("/users", handlers.MakeHTTPHandleFunc(userHandler.Handle))

	log.Println("server running on Port:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
