package models

import "time"

const (
	Type_Teacher = 1
	Type_Student = 2
	Type_Written=00
	Type_Mcq=11
)

type Class struct {
	Name         string   `json:"name,omitempty"`
	Code         string   `json:"code,omitempty"`
	TeacherEmail string   `json:"teacher_email,omitempty"`
	StudentList  []string `json:"student_list,omitempty"`
}

// type Assignment struct {
// 	ID          int          `json:"id,omitempty"`
// 	Class_code  string       `json:"class_code,omitempty"`
// 	Description string       `json:"description,omitempty"`
// 	FileName    string       `json:"file_name,omitempty"`
// 	Submissions []Submission `json:"submissions,omitempty"`
// 	Date        time.Time    `json:"date,omitempty"`
// }

type Assignment struct {
	ID          int       `json:"id,omitempty"`
	Class_code  string    `json:"class_code,omitempty"`
	Type        int       `json:"type,omitempty"`
	Form        Mcq     `json:"form,omitempty"`
	File        Written `json:"file,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Description string    `json:"description,omitempty"`
}

type Mcq struct {
	Ques    []Questions  `json:"ques,omitempty"`
	Answers []string     `json:"answers,omitempty"`
	Soln    []Submission `json:"soln,omitempty"`
}
type Questions struct {
	Question string   `json:"question,omitempty"`
	Options  []string `json:"options,omitempty"`
}

type Written struct {
	FileName    string       `json:"file_name,omitempty"`
	Submissions []Submission `json:"submissions,omitempty"`
}
type Submission struct {
	FileName  string `json:"file_name,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Score     int    `json:"score,omitempty"`
}

type Announcement struct {
	ID          int    `json:"id,omitempty"`
	ClassCode   string `json:"class_code,omitempty"`
	Description string `json:"description,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
}
