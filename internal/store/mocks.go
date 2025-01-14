package store

import (
	"context"
	"database/sql"
)

type MockUsersStore struct {
	db *sql.DB
}

func NewMockUsersStore() Users {
	return &MockUsersStore{
		db: nil,
	}
}

func (s *MockUsersStore) Create(ctx context.Context, u *User) error {
	return nil
}

func (s *MockUsersStore) GetUser(ctx context.Context, u *User) error {
	return nil
}

func (s *MockUsersStore) GetAllUsers(ctx context.Context, u *[]User) error {
	return nil
}
