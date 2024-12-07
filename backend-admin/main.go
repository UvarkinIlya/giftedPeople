package main

import (
	"backend-admin/service"
	"backend-admin/srv"
	"backend-admin/storage"
)

func main() {
	mongoStorage, err := storage.New()
	if err != nil {
		panic(err)
	}

	personService := service.NewPersonService(mongoStorage)
	imageService := service.NewImageService(mongoStorage)

	server := srv.NewServer(personService, imageService, "8080")
	server.Run()
}
