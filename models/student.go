package models

// This file contains the structures/models of data related
// to a Student.

type Student struct {
	Name      string   `json:"name"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	AdmNo     string   `json:"admno"`
	Phone     string   `json:"phone"`
	ClassCode []string `json:"classcode"`
	Type      int      `json:"type"`
}
