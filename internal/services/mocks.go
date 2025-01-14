package services

import (
	"context"

	"github.com/daglamier22/my-clients-be/internal/models"
	"github.com/daglamier22/my-clients-be/internal/store"
)

type MockUsersService struct {
	usersStore store.Storage
}

func NewMockUsersService(usersStore *store.MockUsersStore) Users {
	return &MockUsersService{
		usersStore: usersStore,
	}
}

func (s *MockUsersService) Signup(ctx context.Context, payload models.SignupPayload) (*models.SignupResponse, error) {
	return &models.SignupResponse{
		Status:  "Success",
		Message: "",
		Uuid:    "asdf",
	}, nil
}
