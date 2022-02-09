package models

type Student struct {
	Name string `json:"name"`
	CPG  string `json:"cpg"`
	ID   string `json:"id"`
}

var Students []Student
