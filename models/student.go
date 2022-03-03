package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPG  string `json:"cpg" validate:"len=11, regexp=^[0-9]*$"`
	//ID   string `json:"id"`
}

func ValidateData(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}

	return nil
}
