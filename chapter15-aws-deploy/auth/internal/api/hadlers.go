package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

var (
	tokenSecretKey = os.Getenv("TOKEN_KEY")
)

type GetHashedPassword struct {
	Password string `json:"password" validate:"min=4"`
}

func getHashedPassword(c *gin.Context) {
	var body GetHashedPassword

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": gin.H{"hashedPassword": string(hashedPassword)}, "error": nil})
}

type GetToken struct {
	Password       string `json:"password" validate:"min=4"`
	HashedPassword string `json:"hashedPassword"`
}

func getToken(c *gin.Context) {
	var body GetToken
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(body.HashedPassword), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tokenSecretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": gin.H{"token": tokenString}, "error": nil})
}

type GetTokenConfirmation struct {
	Token string `json:"token"`
}

func getTokenConfirmation(c *gin.Context) {
	var body GetTokenConfirmation
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	if err := verifyToken(body.Token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "OK", "error": nil})
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
