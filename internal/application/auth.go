package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/daglamier22/my-clients-be/internal/models"
	"github.com/daglamier22/my-clients-be/internal/utils"
)

func (app *Application) signupHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.SignupPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := app.UsersService.Signup(r.Context(), payload)
	if err != nil {
		errResp := models.SignupResponse{
			Status:  "Error",
			Message: err.Error(),
		}
		jsonResp, err := json.Marshal(errResp)
		if err != nil {
			log.Fatalf("error handling JSON marshal. Err: %v", err)
		}

		w.Write(jsonResp)
		return
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
