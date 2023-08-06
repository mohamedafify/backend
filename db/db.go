package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	dbConnStr := "postgres://postgres@localhost/backend?sslmode=disable"
	db, err := NewDB(dbConnStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDB(dbConnStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := initializeDB(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initializeDB(db *sql.DB) error {
	dropTable := `DROP TABLE USERS;`
	createTable := `CREATE TABLE IF NOT EXISTS USERS (
		ID UUID PRIMARY KEY,
		FIRST_NAME VARCHAR(100),
		LAST_NAME VARCHAR(100),
		EMAIL VARCHAR(100),
		PHONE_NUMBER VARCHAR(100),
		CREATED_AT TIMESTAMP,
		PASSWORD VARCHAR(400)
	);`

	var query string
	query = dropTable + createTable

	_, err := db.Exec(query)
	return err
}
