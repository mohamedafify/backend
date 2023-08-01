package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	Init() error
	Create(*any) error
	Get(int) (*any, error)
	Update(*any) error
	Delete(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {
	query := `CREATE TABLE IF NOT EXISTS USERS (
		PK SERIAL PRIMARY KEY,
		FIRST_NAME VARCHAR(50),
		LAST_NAME VARCHAR(50),
		PHONENUMBER VARCHAR(20),
		PASSWORD VARCHAR(400)
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) Create(*any) error {
	return nil
}

func (s *PostgresStore) Get(int) (*any, error) {
	return nil, nil
}

func (s *PostgresStore) Update(*any) error {
	return nil
}

func (s *PostgresStore) Delete(int) error {
	return nil
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=backend sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}
