package models

type LoginResponse struct {
	Success bool   `json:"success,omitempty"`
	Token   string `json:"token,omitempty"`
	Error   string  `json:"error,omitempty"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Dash struct {
	Email string `json:"email,omitempty"`
}
