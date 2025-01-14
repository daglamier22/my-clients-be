package application

import (
	// "io"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/daglamier22/my-clients-be/internal/models"
	"github.com/daglamier22/my-clients-be/internal/services"
	"github.com/daglamier22/my-clients-be/internal/store"
	"github.com/stretchr/testify/require"
)

func mockSignup(ctx context.Context, payload models.SignupPayload) (*models.SignupResponse, error) {
	return &models.SignupResponse{
		Status:  "Success",
		Message: "",
		Uuid:    "asdf",
	}, nil
}

func TestSignupHandler(t *testing.T) {
	// app := &Application{
	// 	// Config: cfg,
	// 	// Db:     db,
	// }
	// server := httptest.NewServer(http.HandlerFunc(app.signupHandler))
	// resp, err := http.Post(server.URL, "application/json", )
	// if err != nil {
	// 	t.Error(err)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	t.Errorf("Expected 200 but got %d", resp.StatusCode)
	// }

	// expected := ""
	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Error(err)
	// }
	// if string(b) != expected {
	// 	t.Errorf("expected %s but we got %s", expected, string(b))
	// }
	// Mock dependencies
	mockUsersService := services.NewMockUsersService(store.NewUsersStore(nil))

	app := newTestApplication(t)

	tests := []struct {
		name             string
		requestBody      string
		expectedCode     int
		expectedResponse *models.SignupResponse
	}{
		// Test case 1
		{
			name:         "Happy Path",
			requestBody:  `{"email": "test@example.com", "password": "strong_passw0rd", "username": "test", "first_name": "test", "last_name": "code"}` + "\n",
			expectedCode: http.StatusOK,
		},
		// Test case 2
		{
			name:             "Bad Request - Invalid Payload",
			requestBody:      `{"email": "invalid", "password": 123456}`,
			expectedCode:     http.StatusBadRequest,
			expectedResponse: &models.SignupResponse{Status: "Error", Message: "invalid character 'I' looking for beginning of value"}},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(tc.requestBody))
			w := httptest.NewRecorder()

			app.signupHandler(w, r)

			require.Equal(t, tc.expectedCode, w.Code)

			// Check if expected response was returned
			if tc.expectedResponse != nil {
				var respModel models.SignupResponse
				resp := json.NewDecoder(strings.NewReader(w.Body.String())).Decode(&respModel)
				require.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}
