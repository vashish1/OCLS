package models

const(
	ComputerScience uint = iota
	InformationTechnology
	ElectronicCommunication
	Electrical
	Mechanical
	ElectronicAndElectrical
	Civil
)

type Student struct{
	Uid string
	Name string
	Department uint
	Section uint
	Email string
	MobileNo string
	PassHash []byte
}