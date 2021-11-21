package models

type Teacher struct {
	Name string     `json:"name"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	Type       int      `json:"type"`
	Post       []int    `json:"post"`  //to store announcement ID's
	Class      []string `json:"class"` //to store class ID's
	Assignment []int    `json:"assignment"`      //to store the assignment ID's
}
