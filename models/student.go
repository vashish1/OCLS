package models

// This file contains the structures/models of data related
// to a Student.

type Student struct {
	Name      string   `json:"name,"`
	Password  string   `json:"password,"`
	Email     string   `json:"email,"`
	AdmNo     string   `json:"adm_no,"`
	Phone     string   `json:"phone,"`
	ClassCode []string `json:"class_code"`
	Type      int      `json:"type"`
}
