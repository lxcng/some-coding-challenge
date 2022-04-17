package service

import (
	"fmt"
	"testing"
	"visable/internal/client"
	"visable/internal/config"
	"visable/internal/dto"
	"visable/internal/storage"
)

var (
	addr = "https://employees-api.vercel.app"
	conf = &config.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "password",
		Dbname:   "postgres",
	}
)

func TestCreate(t *testing.T) {
	st, err := storage.NewStorage(conf)
	if err != nil {
		t.Fatal(err)
	}
	srv := NewService(client.NewEmployeeClient(addr), st)
	err = srv.CreateProject(&dto.ProjectCreateReq{
		Name:         "project1",
		Owner:        "84b6dbe4-ee03-4431-b8de-9635f75210f2",
		Participants: []string{"eed4364b-5fc9-4807-8d94-ee114539a5f2", "54e51466-bace-4d91-93d0-b65c8cc1a604"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	st, err := storage.NewStorage(conf)
	if err != nil {
		t.Fatal(err)
	}
	srv := NewService(client.NewEmployeeClient(addr), st)
	prs, err := srv.GetProjects()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prs)
}
