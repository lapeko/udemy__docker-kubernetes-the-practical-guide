package main

import (
	"github.com/lapeko/udemy__docker-kubernetes-the-practical-guide/chapter4/backend/internal/api"
	"log"
)

func main() {
	a := api.New()

	if err := a.SetupStorage(); err != nil {
		log.Fatalln(err)
	}

	a.SetupRoutes()

	if err := a.Serve(); err != nil {
		log.Fatalln(err)
	}

	defer a.Disconnect()
}
