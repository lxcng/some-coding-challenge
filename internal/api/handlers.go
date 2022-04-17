package api

import (
	"encoding/json"
	"strconv"
	"visable/internal/dto"

	fiber "github.com/gofiber/fiber/v2"
)

func (x *Api) getProjects(c *fiber.Ctx) error {
	prs, err := x.storage.GetProjects()
	if err != nil {
		return err
	}
	return c.JSON(prs)
}

func (x *Api) createProject(c *fiber.Ctx) error {
	req := &dto.ProjectCreateReq{}
	err := json.Unmarshal(c.Body(), req)
	if err != nil {
		return err
	}
	err = x.storage.CreateProject(req)
	if err != nil {
		return err
	}
	return nil
}

func (x *Api) updateProject(c *fiber.Ctx) error {
	req := &dto.ProjectUpdateReq{}
	err := json.Unmarshal(c.Body(), req)
	idRaw := c.Params("id")
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return err
	}
	req.ID = uint(id)
	err = x.storage.UpdateProject(req)
	if err != nil {
		return err
	}
	return nil
}
