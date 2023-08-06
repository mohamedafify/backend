package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/mohamedafify/backend/handlers"
	"github.com/mohamedafify/backend/services"
)

func InitUsers(db *sql.DB, router *mux.Router) {
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	router.HandleFunc("/users", handlers.MakeHTTPHandleFunc(userHandler.Handle))
}
