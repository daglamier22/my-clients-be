package main

import (
	"fmt"
	"os"

	"github.com/daglamier22/my-clients-be/internal/server"
)

func main() {

	server := server.NewServer()

	fmt.Printf("Server is running at %v\n\r", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
