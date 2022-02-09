package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	CPG  string `json:"cpg"`
	ID   string `json:"id"`
}

var Students []Student
