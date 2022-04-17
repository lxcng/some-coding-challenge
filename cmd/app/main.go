package main

import (
	"visable/internal/api"
	"visable/internal/client"
	"visable/internal/config"
	"visable/internal/service"
	"visable/internal/storage"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	storage, err := storage.NewStorage(conf)
	if err != nil {
		panic(err)
	}
	err = storage.Migrate()
	if err != nil {
		panic(err)
	}
	service := service.NewService(client.NewEmployeeClient(conf.ClientAddr), storage)
	server, err := api.NewApi(service)
	if err != nil {
		panic(err)
	}
	if err := server.Serve(conf.ServeAddr); err != nil {
		panic(err)
	}
}
