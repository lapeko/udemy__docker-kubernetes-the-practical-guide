package main

import (
	"github.com/lapeko/udemy__docker-kubernetes-the-practical-guide/chapter15-aws-deploy/users/internal/api"
	"log"
)

var port = 3000

func main() {
	a := api.New()
	if err := a.ConnectStorage(); err != nil {
		log.Fatalf("Storage connection failed: %v", err)
	}
	log.Fatalln(a.Serve(port))
}
