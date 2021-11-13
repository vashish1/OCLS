package models

import "time"

type Announcment struct {
}

type Assignment struct {
	ID          int          `json:"id"`
	Description string       `json:"description"`
	FileName    string       `json:"file_name"`
	Submissions []Submission `json:"submissions"`
}

type Submission struct {
	FileName  string    `json:"file_name"`
	Email     string    `json:"email"`
	Timestamp time.Time `json:"timestamp"`
}
