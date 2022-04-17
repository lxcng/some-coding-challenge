package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name         string
	State        ProjectState `gorm:"default:planned"`
	Progress     float64
	Owner        *Owner         `gorm:"foreignKey:ProjectID"`
	Participants []*Participant `gorm:"foreignKey:ProjectID"`
}
