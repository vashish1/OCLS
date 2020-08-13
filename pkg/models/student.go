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
	Uid        string `json:"uid,omitempty"`
	Name       string `json:"name,omitempty"`
	Department string   `json:"department,omitempty"`
	Section    uint   `json:"section,omitempty"`
	Email      string `json:"email,omitempty"`
	MobileNo   string `json:"mobile_no,omitempty"`
	PassHash   string `json:"pass_hash,omitempty"`
	Freeze     bool   `json:"freeze,omitempty"`
}
