package services

import (
	"database/sql"

	"github.com/mohamedafify/backend/models"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	query := `INSERT INTO users (id, first_Name, Last_Name, email, password, phone_Number, created_At) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNumber, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
