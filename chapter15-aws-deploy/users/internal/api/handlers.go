package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
)

var authUrl = func() string {
	url := os.Getenv("AUTH_URL")
	if url == "" {
		panic("AUTH_URL is not set")
	}
	return url
}()

type CreateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

func createUser(c *gin.Context) {
	var body CreateUser

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}

	s := getApiStorage()
	user, err := s.Users.GetByEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": err.Error()})
		return
	}

	if user != nil {
		c.JSON(http.StatusConflict, gin.H{"response": nil, "error": fmt.Sprintf("Email %s already taken", body.Email)})
		return
	}

	jsonData, err := json.Marshal(map[string]string{"password": body.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": err.Error()})
		return
	}

	resp, err := http.Post(fmt.Sprintf("%s/hashed-pw", authUrl), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": err.Error()})
		return
	}

	defer resp.Body.Close()

	var postResponse struct {
		Error    string `json:"error"`
		Response struct {
			HashedPassword string `json:"hashedPassword"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&postResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": "Auth request parse failure"})
		return
	}

	if postResponse.Error != "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": postResponse.Error})
		return
	}

	if postResponse.Response.HashedPassword == "" {
		c.JSON(http.StatusBadGateway, gin.H{"response": nil, "error": "Get hashed password error"})
		return
	}

	_, err = s.Users.Create(body.Email, postResponse.Response.HashedPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"response": "OK", "error": nil})
}

type VerifyUserBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
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

	s := getApiStorage()
	user, err := s.Users.GetByEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"response": nil, "error": fmt.Sprintf("User with email %s not exists", body.Email)})
		return
	}

	jsonData, err := json.Marshal(map[string]string{"password": body.Password, "hashedPassword": user.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": err.Error()})
		return
	}

	resp, err := http.Post(fmt.Sprintf("%s/token", authUrl), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": err.Error()})
		return
	}

	defer resp.Body.Close()

	var postResponse struct {
		Error    string `json:"error"`
		Response struct {
			Token string `json:"token"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&postResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": "Auth request parse failure"})
		return
	}

	if postResponse.Error != "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": postResponse.Error})
		return
	}

	if postResponse.Response.Token == "" {
		c.JSON(http.StatusBadGateway, gin.H{"response": nil, "error": "Get token error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": gin.H{"token": postResponse.Response.Token}, "error": nil})
}
