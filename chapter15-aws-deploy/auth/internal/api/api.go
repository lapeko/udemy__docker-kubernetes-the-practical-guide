package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.Engine
}

var api *Api

func NewApi() *Api {
	if api == nil {
		api = &Api{
			router: gin.Default(),
		}
	}
	return api
}

func (a *Api) Run(port int) error {
	a.router.POST("/hashed-pw", getHashedPassword)
	a.router.POST("/token", getToken)
	a.router.POST("/verify-token", getTokenConfirmation)

	return a.router.Run(fmt.Sprintf(":%d", port))
}
