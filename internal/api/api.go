package api

import (
	fiber "github.com/gofiber/fiber/v2"
)

type Api struct {
	app     *fiber.App
	storage Service
}

func NewApi(storage Service) (*Api, error) {
	app := fiber.New(fiber.Config{ServerHeader: "Some Bank"})
	res := &Api{
		app:     app,
		storage: storage,
	}
	return res, nil
}

func (x *Api) Serve(addr string) error {
	x.app.Get("/api/project", x.getProjects)
	x.app.Post("/api/project", x.createProject)
	x.app.Post("/api/project/:id", x.updateProject)
	return x.app.Listen(addr)
}
