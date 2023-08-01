package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mohamedafify/backend/storage"
	"github.com/mohamedafify/backend/types"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	log.Println("server running on Port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetUser(w, r)
	case "POST":
		return s.handleCreateUser(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	user := types.NewUser("Mohamed", "Afify", "+201000000000", "Password@12345")
	return WriteJSON(w, http.StatusOK, user)
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleSignup(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
