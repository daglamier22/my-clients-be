package services

import (
	"context"
	"log"

	"github.com/daglamier22/my-clients-be/internal/models"
	"github.com/daglamier22/my-clients-be/internal/store"
)

type UsersService struct {
	usersStore store.UsersStore
}

func NewUsersService(usersStore store.UsersStore) UsersService {
	return UsersService{
		usersStore: usersStore,
	}
}

func (s *UsersService) Signup(ctx context.Context, payload models.SignupPayload) (*models.SignupResponse, error) {
	u := &store.User{
		Username:  payload.Username,
		Password:  payload.Password,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}

	err := s.usersStore.Create(ctx, u)
	log.Printf("id: %v err: %v", u.Id, err)
	if err != nil {
		return nil, err
	}

	return &models.SignupResponse{
		Status:  "Success",
		Message: "",
		Uuid:    "asdf",
	}, nil
}
