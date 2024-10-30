package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/public", "./public")
	r.Static("/feedback", "./permanent-data")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Feedback app"})
	})

	r.GET("/exists", func(c *gin.Context) {
		c.HTML(http.StatusOK, "exists.tmpl", gin.H{})
	})

	r.POST("/", func(c *gin.Context) {
		title := c.PostForm("title")
		text := c.PostForm("text")

		fileName := strings.ToLower(title)
		tempFilePath := fmt.Sprintf("./temporary-data/%s", fileName)
		permFilePath := fmt.Sprintf("./permanent-data/%s", fileName)

		file, err := os.Create(tempFilePath)

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		_, err = file.Write([]byte(text))
		if err != nil {
			log.Fatalln(err)
		}

		if _, err := os.Stat(permFilePath); err == nil {
			c.Redirect(http.StatusSeeOther, "/exists")
			return
		}

		if err := os.Rename(tempFilePath, permFilePath); err != nil {
			log.Fatalln(err)
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":3000")
}
