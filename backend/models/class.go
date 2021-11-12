package models

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
