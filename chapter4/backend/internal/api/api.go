package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lapeko/udemy__docker-kubernetes-the-practical-guide/chapter4/backend/storage"
	"os"
)

type Api struct {
	storage   *storage.Storage
	ginEngine *gin.Engine
}

var singleApi *Api

func New() *Api {
	if singleApi == nil {
		singleApi = &Api{}
	}
	return singleApi
}

func (a *Api) SetupRoutes() {
	r := gin.Default()
	a.ginEngine = r
	a.ginEngine.Use(corsMiddleware())

	r.GET("/", a.getHandler)
	r.POST("/", a.postHandler)
	r.DELETE("/:id", a.deleteHandler)
}

func (a *Api) SetupStorage() error {
	a.storage = storage.New()
	return a.storage.Connect()
}

func (a *Api) Serve() error {
	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("PORT environment variable not set")
	}

	return a.ginEngine.Run(fmt.Sprintf(":%s", port))
}

func (a *Api) Disconnect() {
	a.storage.Disconnect()
}
