package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func createUser(c *gin.Context) {

}

type VerifyUserBody struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=4"`
}

func verifyUser(c *gin.Context) {
	var body VerifyUserBody

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

}
