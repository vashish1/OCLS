package models

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
