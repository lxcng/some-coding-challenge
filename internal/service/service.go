package service

import (
	"errors"
	"fmt"
	"visable/internal/client"
	"visable/internal/dto"
	"visable/internal/model"

	"gorm.io/gorm"
)

var (
	ErrInvalidState    = errors.New("invalid project state")
	ErrInvalidProgress = errors.New("invalid project progress")
	ErrEmptyOwner      = errors.New("empty owner")
	ErrNonmanagerOwner = errors.New("nonmanager owner")
)

type Service struct {
	client  EmployeeClient
	storage Storage
}

func NewService(client EmployeeClient, storage Storage) *Service {
	return &Service{
		client:  client,
		storage: storage,
	}
}

func (x *Service) GetProjects() ([]*dto.Project, error) {
	projects, err := x.storage.GetProjects()
	if err != nil {
		return nil, err
	}
	res := make([]*dto.Project, 0, len(projects))
	for _, pr := range projects {
		project := &dto.Project{
			ID:       pr.ID,
			Name:     pr.Name,
			State:    string(pr.State),
			Progress: pr.Progress,
			Owner: &dto.Employee{
				Id:         pr.Owner.EmployeeId,
				FirstName:  pr.Owner.FirstName,
				LastName:   pr.Owner.LastName,
				Department: pr.Owner.Department,
			},
		}
		parts := make([]*dto.Employee, 0, len(pr.Participants))
		for _, p := range pr.Participants {
			parts = append(parts, &dto.Employee{
				Id:        p.EmployeeId,
				FirstName: p.FirstName,
				LastName:  p.LastName,
			})
		}
		project.Participants = parts
		res = append(res, project)
	}
	return res, nil
}

func (x *Service) CreateProject(req *dto.ProjectCreateReq) error {
	prState, ok := model.ValidateProjectState(req.State)
	if !ok {
		return ErrInvalidState
	}
	if req.Progress < 0 || req.Progress > 100 {
		return ErrInvalidProgress
	}
	if req.Owner == "" {
		return ErrEmptyOwner
	}
	owner, err := x.client.GetEmployeeById(req.Owner)
	if err != nil {
		return err
	}
	if owner.Role != client.RoleManager {
		return ErrNonmanagerOwner
	}
	project := &model.Project{
		Name:     req.Name,
		State:    prState,
		Progress: req.Progress,
		Owner: &model.Owner{
			EmployeeId: owner.Id,
			FirstName:  owner.FirstName,
			LastName:   owner.LastName,
			Department: string(owner.Department),
		},
	}
	if len(req.Participants) > 0 {
		parts, errs := x.client.GetEmployeesByIds(req.Participants)
		if len(errs) > 0 {
			return fmt.Errorf("failed to fetch participants: %v", errs)
		}
		participants := make([]*model.Participant, 0, len(parts))
		for _, part := range parts {
			if part.Department != owner.Department {
				return fmt.Errorf("participant %v is in the wrong department", part.Id)
			}
			participants = append(participants, &model.Participant{
				EmployeeId: part.Id,
				FirstName:  part.FirstName,
				LastName:   part.LastName,
			})
		}
		project.Participants = participants
	}
	return x.storage.CreateProject(project)
}

func (x *Service) UpdateProject(req *dto.ProjectUpdateReq) error {
	prState, ok := model.ValidateProjectState(req.State)
	if !ok {
		return ErrInvalidState
	}
	if req.Progress < 0 || req.Progress > 100 {
		return ErrInvalidProgress
	}
	project := &model.Project{
		Model: gorm.Model{
			ID: req.ID,
		},
		Name:     req.Name,
		State:    prState,
		Progress: req.Progress,
	}
	return x.storage.UpdateProject(project)
}
