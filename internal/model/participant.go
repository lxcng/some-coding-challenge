package model

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	ProjectID  uint
	EmployeeId string
	FirstName  string
	LastName   string
}
