package models

import "time"

const (
	Type_Teacher = 1
	Type_Student = 2
)

type Class struct {
	Name         string   `json:"name,omitempty"`
	Code         string   `json:"code,omitempty"`
	TeacherEmail string   `json:"teacher_email,omitempty"`
	StudentList  []string `json:"student_list,omitempty"`
}

type Assignment struct {
	ID          int          `json:"id"`
	Description string       `json:"description"`
	FileName    string       `json:"file_name"`
	Submissions []Submission `json:"submissions"`
	Date        time.Time    `json:"date"`
}

type Submission struct {
	FileName  string    `json:"file_name"`
	Email     string    `json:"email"`
	Timestamp time.Time `json:"timestamp"`
}

type Announcement struct{
	ID int
	ClassCode string
	Description string
	Timestamp string
}