package models

// This file contains the structures/models of data related
// to a Student.

type Student struct {
	Name      string   `json:"name,omitempty"`
	Email     string   `json:"email,omitempty"`
	AdmNo     string   `json:"adm_no,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	ClassCode []string `json:"class_code"`
	Type      int      `json:"type"`
}
