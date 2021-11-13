package models

type Teacher struct {
	Email string        `json:"email,omitempty"`
	Type  int           `json:"type,omitempty"`
	Post  []int `json:"post,omitempty"` //to store announcement ID's
	Class []string      `json:"class,omitempty"` //to store class ID's
	Assignment []int  `json:"assignment"` //to store the assignment ID's
}
