package model

import "gorm.io/gorm"

type Owner struct {
	gorm.Model
	ProjectID  uint
	EmployeeId string
	FirstName  string
	LastName   string
	Department string
}
