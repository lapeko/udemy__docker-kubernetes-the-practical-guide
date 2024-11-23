package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var PORT = os.Getenv("PORT")

func main() {
	fmt.Println("PORT", PORT)
	r := gin.Default()

	r.GET("/verify-token/:token", func(c *gin.Context) {
		token, ok := c.Params.Get("token")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token is not defined"})
			return
		}
		if token != "abc" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Provided token is invalid or expired"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Valid token", "uid": "u1"})
	})

	r.GET("/token/:hashedPassword/:enteredPassword", func(c *gin.Context) {
		hashed, ok := c.Params.Get("hashedPassword")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Hashed password is not defined"})
			return
		}
		entered, ok := c.Params.Get("enteredPassword")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Entered password is not defined"})
			return
		}
		if hashed != fmt.Sprintf("%s_hash", entered) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
			return
		}
		token := "abc"
		c.JSON(http.StatusOK, gin.H{"message": "Token created", "token": token})
	})

	r.GET("/hashed-password/:password", func(c *gin.Context) {
		password, ok := c.Params.Get("password")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Entered password is not defined"})
			return
		}
		hashedPassword := fmt.Sprintf("%s_hash", password)
		c.JSON(http.StatusOK, gin.H{"hashedPassword": hashedPassword})
	})

	fmt.Printf("Server is running on port: %s\n", PORT)
	log.Fatalln(r.Run(fmt.Sprintf(":%s", PORT)))
}
