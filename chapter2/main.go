package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/public", "./public")

	goal := ""

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, fmt.Sprintf(`
			<!doctype html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport"
					  content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<link rel="stylesheet" href="/public/style.css">
				<title>Document</title>
			</head>
			<body>
				<p>%s</p>
				<form method="post" action="/">
					<input name="goal">
					<button>Set goal</button>
				</form>
			</body>
			</html>
		`, goal))
	})

	r.POST("/", func(c *gin.Context) {
		goal = c.PostForm("goal")
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":3000")
}
