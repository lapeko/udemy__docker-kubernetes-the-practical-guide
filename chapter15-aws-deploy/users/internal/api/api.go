package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.Engine
}

var api *Api

func (a *Api) NewApi() *Api {
	if api == nil {
		api = &Api{
			router: gin.Default(),
		}
	}
	return api
}

func (a *Api) Run(port int) error {
	a.router.POST("/signup", createUser)
	a.router.POST("/login", verifyUser)
	return a.router.Run(fmt.Sprintf(":%d", port))
}
