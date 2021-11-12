package models

type Teacher struct {
	Email string        `json:"email,omitempty"`
	Type  int           `json:"type,omitempty"`
	Post  []Announcment `json:"post,omitempty"`
	Class []string      `json:"class,omitempty"`
}
