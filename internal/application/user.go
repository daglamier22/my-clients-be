package application

import (
	// "encoding/json"
	// "fmt"
	// "log"
	"net/http"

	"github.com/daglamier22/my-clients-be/internal/store"
	"github.com/daglamier22/my-clients-be/internal/utils"
)

// func (app *Application) testGetUser(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	user, err := s.db.GetUser(id)
// 	fmt.Print(err)
// 	if err != nil {
// 		log.Print(err)
// 		resp := make(map[string]string)
// 		resp["error"] = "ouch"

// 		jsonResp, err := json.Marshal(resp)
// 		if err != nil {
// 			log.Fatalf("error handling JSON marshal. Err: %v", err)
// 		}
// 		fmt.Printf("resp: %v", err)

// 		_, _ = w.Write(jsonResp)
// 		return
// 	}

// 	resp := make(map[string]string)
// 	resp["username"] = user.Username
// 	resp["password"] = user.Password

// 	jsonResp, err := json.Marshal(resp)
// 	if err != nil {
// 		log.Fatalf("error handling JSON marshal. Err: %v", err)
// 	}

// 	_, _ = w.Write(jsonResp)
// }

func (app *Application) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []store.User{}
	// err := app.Store.Users.GetAllUsers(r.Context(), &users)
	// // users, err := a.db.GetAllUsers()
	// if err != nil {
	// 	log.Print(err)
	// 	resp := make(map[string]string)
	// 	resp["error"] = "ouch"

	// 	jsonResp, err := json.Marshal(resp)
	// 	if err != nil {
	// 		log.Fatalf("error handling JSON marshal. Err: %v", err)
	// 	}
	// 	fmt.Printf("resp: %v", err)

	// 	_, _ = w.Write(jsonResp)
	// 	return
	// }

	// resp := make(map[string]string)
	// resp["username"] = users[0].Username
	// resp["password"] = users[0].Password
	// resp["users"], err = json.Marshal(users)

	// jsonResp, err := json.Marshal(users)//resp)
	// if err != nil {
	// 	log.Fatalf("error handling JSON marshal. Err: %v", err)
	// }

	// _, _ = w.Write(jsonResp)
	utils.WriteJSON(w, http.StatusOK, users)
}
