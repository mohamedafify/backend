package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/mohamedafify/backend/models"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

/*
func (uh *UserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	makeHTTPHandleFunc(uh.handle)
}
*/

func (uh *UserHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodPost:
		return uh.CreateUser(w, r)
	}
	return fmt.Errorf("%s Method is not allowed", r.Method)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {

	var req models.CreateUserRequest
	if err := DecodeBody(r, &req); err != nil {
		return err
	}

	user, err := models.CreateUser(req.FirstName, req.LastName, req.PhoneNumber, req.Email, req.Password)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (id, firstName, LastName, email, password, phoneNumber, createdAt) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, insertErr := uh.db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNumber, user.CreatedAt)
	if insertErr != nil {
		return insertErr
	}

	return WriteJSON(w, http.StatusOK, user)
}

/*
func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	users := []models.User{}
	rows, err := uh.db.Query("SELECT id, name, email, location FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

	return nil
}
*/
