package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	PORT       = os.Getenv("PORT")
	FS_DB_PATH = os.Getenv("FS_DB_PATH")
	AUTH_PATH  = os.Getenv("AUTH_PATH")
)

type CreateTaskBody struct {
	Task string `json:"task"`
}

type VerifyBody struct {
	Uid string `json:"uid"`
}

func main() {
	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		_, err := extractAndVerifyUid(c.GetHeader("Authorization"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		content, err := os.ReadFile(FS_DB_PATH)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		tasks := strings.Split(string(content), "\n")

		if len(tasks) == 0 {
			c.JSON(http.StatusOK, gin.H{"tasks": []interface{}{}})
		}

		c.JSON(http.StatusOK, gin.H{"tasks": tasks[:len(tasks)-1]})
	})

	r.POST("/tasks", func(c *gin.Context) {
		_, err := extractAndVerifyUid(c.GetHeader("Authorization"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var taskBody CreateTaskBody
		if err := c.ShouldBind(&taskBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		file, err := os.OpenFile(FS_DB_PATH, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer file.Close()

		if _, err := file.Write([]byte(fmt.Sprintf("%s\n", taskBody.Task))); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Task stored", "createdTask": taskBody.Task})
	})

	log.Fatalln(r.Run(fmt.Sprintf(":%s", PORT)))
}

func extractAndVerifyUid(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("no token provided")
	}

	parts := strings.Split(authorization, "Bearer ")

	if len(parts) != 2 {
		return "", errors.New("no token provided")
	}

	token := parts[1]

	res, err := http.Get(fmt.Sprintf("%s/verify-token/%s", AUTH_PATH, token))

	if err != nil {
		return "", err
	}

	var verifyBody VerifyBody
	if err := json.NewDecoder(res.Body).Decode(&verifyBody); err != nil {
		return "", err
	}

	return verifyBody.Uid, nil
}
