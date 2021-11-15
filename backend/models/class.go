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
	ID          int          `json:"id,omitempty"`
	Class_code  string       `json:"class_code,omitempty"`
	Description string       `json:"description,omitempty"`
	FileName    string       `json:"file_name,omitempty"`
	Submissions []Submission `json:"submissions,omitempty"`
	Date        time.Time    `json:"date,omitempty"`
}

type Submission struct {
	FileName  string    `json:"file_name"`
	Email     string    `json:"email"`
	Timestamp time.Time `json:"timestamp"`
}

type Announcement struct {
	ID          int    `json:"id,omitempty"`
	ClassCode   string `json:"class_code,omitempty"`
	Description string `json:"description,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
}
