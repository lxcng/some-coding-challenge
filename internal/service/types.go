package service

import (
	"visable/internal/client"
	"visable/internal/model"
)

type EmployeeClient interface {
	GetEmployeeById(id string) (*client.Employee, error)
	GetEmployeesByIds(ids []string) ([]*client.Employee, []error)
}

type Storage interface {
	GetProjects() ([]*model.Project, error)
	CreateProject(project *model.Project) error
	UpdateProject(project *model.Project) error
}
