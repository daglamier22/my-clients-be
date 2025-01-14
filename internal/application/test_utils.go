package application

import (
	"testing"

	"github.com/daglamier22/my-clients-be/internal/services"
	"github.com/daglamier22/my-clients-be/internal/store"
)

func newTestApplication(t *testing.T) *Application {
	t.Helper()

	mockStore := store.NewMockUsersStore()
	mockService := services.NewMockUsersService(&mockStore)

	return &Application{
		UsersService: mockService,
	}
	// return nil
}
