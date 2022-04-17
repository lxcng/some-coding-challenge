package api

import "visable/internal/dto"

type Service interface {
	GetProjects() ([]*dto.Project, error)
	CreateProject(req *dto.ProjectCreateReq) error
	UpdateProject(req *dto.ProjectUpdateReq) error
}
