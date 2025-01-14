package store

import (
	"context"
	"database/sql"
)

type Users interface {
	Create(ctx context.Context, u *User) error
	GetUser(ctx context.Context, u *User) error
	GetAllUsers(ctx context.Context, u *[]User) error
}

type Storage struct {
	Users Users
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
	}
}
