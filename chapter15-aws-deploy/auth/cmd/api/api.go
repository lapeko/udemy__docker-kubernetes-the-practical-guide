package main

import (
	"github.com/lapeko/udemy__docker-kubernetes-the-practical-guide/chapter15-aws-deploy/auth/internal/api"
	"log"
)

var (
	PORT = 3000
)

func main() {
	a := api.NewApi()
	log.Fatalln(a.Run(PORT))
}
