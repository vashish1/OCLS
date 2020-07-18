package models

type LoginResponse struct {
	Success bool   `json:"success,omitempty"`
	Token   string `json:"token,omitempty"`
}

type Number struct{
	
}