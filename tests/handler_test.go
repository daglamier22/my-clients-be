package tests

import (
	"fmt"
	"net/http"

	// "net/http/httptest"
	"testing"

	"github.com/daglamier22/my-clients-be/internal/application"
	"github.com/daglamier22/my-clients-be/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

// func TestHandler(t *testing.T) {
// 	s := &server.Server{}
// 	server := httptest.NewServer(http.HandlerFunc(s.HelloWorldHandler))
// 	defer server.Close()
// 	resp, err := http.Get(server.URL)
// 	if err != nil {
// 		t.Fatalf("error making request to server. Err: %v", err)
// 	}
// 	defer resp.Body.Close()
// 	// Assertions
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", resp.Status)
// 	}
// 	expected := "{\"message\":\"Hello World\"}"
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("error reading response body. Err: %v", err)
// 	}
// 	if expected != string(body) {
// 		t.Errorf("expected response body to be %v; got %v", expected, string(body))
// 	}
// }

func TestHealthHandler(t *testing.T) {
	cfg := application.Config{
		Addr: fmt.Sprintf(":%d", 8080),
	}
	app := &application.Application{
		Config: cfg,
		Db:     database.NewTest(),
	}
	app.NewApplication()
	defer app.Server.Close()
	resp, err := http.Get("http://" + app.Server.Addr + "/health")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
}
