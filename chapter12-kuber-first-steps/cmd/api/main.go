package main

import "github.com/lapeko/udemy__docker-kubernetes-the-practical-guide/chapter12-kuber-first-steps/internal/api"

func main() {
	a := api.New()
	a.Start()
}
