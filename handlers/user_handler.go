package handlers

import (
	"fmt"
	"net/http"

	"github.com/mohamedafify/backend/models"
	"github.com/mohamedafify/backend/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (uh *UserHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodPost:
		return uh.CreateUser(w, r)
	}
	return fmt.Errorf("%s method is not allowed", r.Method)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {

	var req models.CreateUserRequest
	if err := GetBody(r, &req); err != nil {
		return err
	}

	user, err := models.CreateUser(req.FirstName, req.LastName, req.PhoneNumber, req.Email, req.Password)
	if err != nil {
		return err
	}

	if dbErr := uh.service.CreateUser(user); dbErr != nil {
		return dbErr
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
