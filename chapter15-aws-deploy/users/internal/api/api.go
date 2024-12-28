package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lapeko/udemy__docker-kubernetes-the-practical-guide/chapter15-aws-deploy/users/internal/storage"
	"log"
)

type Api struct {
	router  *gin.Engine
	storage *storage.Storage
}

var api *Api

func New() *Api {
	if api == nil {
		api = &Api{
			router: gin.Default(),
		}
	} else {
		log.Println("API instance already exists. Reusing the existing instance.")
	}
	return api
}

func (a *Api) ConnectStorage() error {
	if a.storage == nil {
		s := storage.NewStorage()
		if err := s.Connect(); err != nil {
			return fmt.Errorf("Api.ConnectStorage error: %w", err)
		}
		a.storage = s
	} else {
		log.Println("Storage already connected. Reusing the existing instance.")
	}
	return nil
}

func (a *Api) Serve(port int) error {
	if a.storage == nil {
		return errors.New("Api.Storage is not connected")
	}
	a.router.POST("/signup", createUser)
	a.router.POST("/login", verifyUser)
	return a.router.Run(fmt.Sprintf(":%d", port))
}
