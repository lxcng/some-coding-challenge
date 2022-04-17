package storage

import "visable/internal/model"

func (x *Storage) GetProjects() ([]*model.Project, error) {
	res := make([]*model.Project, 0)
	return res, x.db.Model(new(model.Project)).Preload("Owner").Preload("Participants").Find(&res).Error
}

func (x *Storage) CreateProject(project *model.Project) error {
	return x.db.Create(project).Error
}

func (x *Storage) UpdateProject(project *model.Project) error {
	return x.db.Model(project).Updates(&model.Project{Name: project.Name, State: project.State, Progress: project.Progress}).Error
}
