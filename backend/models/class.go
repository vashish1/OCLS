package models

const  (
	Type_Teacher = 1
	Type_Student = 2
)
type Class struct{
	Name string
	TeacherEmail string
	StudentList  []Student
}