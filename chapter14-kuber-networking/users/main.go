package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	PORT     = os.Getenv("PORT")
	AUTH_URL = os.Getenv("AUTH_URL")
)

type RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthTokenBody struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func main() {
	if PORT == "" {
		log.Fatalln("Environment PORT is not provided")
	}
	if AUTH_URL == "" {
		log.Fatalln("Environment AUTH_URL is not provided")
	}

	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		var body RequestBody

		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email := strings.Trim(body.Email, " ")
		password := strings.Trim(body.Password, " ")

		if email == "" || password == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Provided email or password is empty"})
			return
		}

		res, err := http.Get(fmt.Sprintf("%s/hashed-password/%s", AUTH_URL, password))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Println("email:", email, "hashedPassword:", res)
		c.JSON(http.StatusCreated, gin.H{"message": "User created!"})
	})

	r.POST("/login", func(c *gin.Context) {
		var body RequestBody

		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email := strings.Trim(body.Email, " ")
		password := strings.Trim(body.Password, " ")

		if email == "" || password == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Provided email or password is empty"})
			return
		}

		hashedPassword := fmt.Sprintf("%s_hash", password)
		res, err := http.Get(fmt.Sprintf("%s/token/%s/%s", AUTH_URL, hashedPassword, password))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if res.StatusCode != http.StatusOK {
			c.JSON(res.StatusCode, gin.H{"error": "Logging in failed!"})
			return
		}

		var resBody AuthTokenBody

		if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": resBody.Token})
	})

	log.Printf("Server is running on port \":%s\"\n", PORT)
	log.Fatalln(r.Run(fmt.Sprintf(":%s", PORT)))
}
