package services

import (
	"context"

	"github.com/daglamier22/my-clients-be/internal/models"
	"github.com/daglamier22/my-clients-be/internal/store"
)

type Users interface {
	Signup(ctx context.Context, payload models.SignupPayload) (*models.SignupResponse, error)
}

type Service struct {
	Users Users
}

func NewServices(store store.Storage) Service {
	return Service{
		Users: &UsersService{store.Users},
	}
}
