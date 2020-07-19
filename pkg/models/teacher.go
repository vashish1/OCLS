package models

type Teacher struct {
	Uid        string `json:"uid,omitempty"`
	Department string `json:"department,omitempty"`
	Name       string `json:"name,omitempty"`
}
