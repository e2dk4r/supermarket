package cockroachdb

import (
	"context"

	"github.com/e2dk4r/supermarket"
	"github.com/jackc/pgx/v4"
)

type UserService struct {
	PasswordService supermarket.PasswordService
	Conn            *pgx.Conn
}

func (us *UserService) User(username string) (*supermarket.User, error) {
	var id string
	var password string

	row := us.Conn.QueryRow(context.Background(), "SELECT id, password FROM users WHERE username = $1", username)
	err := row.Scan(&id, &password)
	if err != nil {
		return nil, err
	}

	return &supermarket.User{
		Id:       id,
		Username: username,
		Password: password,
	}, nil
}

func (us *UserService) CreateUser(user *supermarket.User) error {
	hash, err := us.PasswordService.Hash(user.Password)
	if err != nil {
		return err
	}

	var id string
	row := us.Conn.QueryRow(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", user.Username, hash)
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	user.Id = id
	return nil
}
