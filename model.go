package model;

type Student struct {
	ID  int `json:"id"`
	Name string `json:"name"`
	Age int `json: "age"`
	Subject string `json: "subject"`
	Class string `json: "class"`

}