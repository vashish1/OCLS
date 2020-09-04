package models

const(
	CSE string = "Computer Science Engineering"
	IT = "Information Technology"
	ECE= "Electronic Engineering"
	EE= "Electrical Engineering"
	ME= "Mechanical Engineering"
	EEE= "Electrical & Electronic Engineering"
	CE= "Civil Engineering"
)

type Student struct {
	Uid        string `json:"uid"`
	Name       string `json:"name"`
	Department string   `json:"department"`
	Section    uint   `json:"section"`
	Email      string `json:"email"`
	MobileNo   string `json:"mobile_no"`
	PassHash   string `json:"passhash,omitempty"`
	Freeze     bool   `json:"freeze"`
}
