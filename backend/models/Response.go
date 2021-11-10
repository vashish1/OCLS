package models

type Response struct{
	Success bool
	Error string
	Message string
    Data interface{}
}
