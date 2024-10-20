package models

type SignupPayload struct {
	Username  string `json:"username" validate:"required,max=300"`
	Password  string `json:"password" validate:"required,min=6,max=300"`
	FirstName string `json:"first_name" validate:"required,max=300"`
	LastName  string `json:"last_name" validate:"required,max=300"`
	Email     string `json:"email" validate:"required,email,max=300"`
}

type SignupResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}
