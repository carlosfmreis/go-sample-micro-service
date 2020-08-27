package user

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// Repository interface
type Repository interface {
	CreateUser(context context.Context, user User) error
	GetUser(context context.Context, id int) (User, error)
}

func (repository *repository) CreateUser(context context.Context, user User) error {
	sql := "INSERT INTO users (email, password) VALUES (?, ?);"

	_, err := repository.db.ExecContext(context, sql, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repository *repository) GetUser(context context.Context, id int) (User, error) {
	var user User

	err := repository.db.QueryRow("SELECT * FROM users WHERE id = ?;", id).Scan(&user.ID, &user.Email, &user.Password)

	return user, err
}

// NewRepository method
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
