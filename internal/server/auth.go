package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Body struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}

func (b *Body) Bind(r *http.Request) error {
	return nil
}

func (s *Server) SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Body: %v\n", r.Body)
	
	body := &Body{}
	// var body Body

	err := json.NewDecoder(r.Body).Decode(&body)
	// // err := render.Bind(r, body)
	if err != nil {
		fmt.Printf("err2: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// username := body.Username
	fmt.Printf("hi")// %v\n", username)
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
