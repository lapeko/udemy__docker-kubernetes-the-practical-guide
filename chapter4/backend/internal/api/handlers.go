package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lapeko/udemy__docker-kubernetes-the-practical-guide/chapter4/backend/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type PostTodoRequest struct {
	Title string `json:"title"`
}

func (a *Api) getHandler(c *gin.Context) {
	todos, err := a.storage.GetAllTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"OK": false, "payload": nil, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"OK": true, "payload": gin.H{"todos": todos}, "error": nil})
}

func (a *Api) postHandler(c *gin.Context) {
	todo := &PostTodoRequest{}

	if err := c.BindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"OK": false, "payload": nil, "error": errors.New("wrong request format")})
		return
	}

	response, err := a.storage.InsertTodo(&storage.Todo{
		Title: todo.Title,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"OK": false, "payload": nil, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": true, "payload": response, "error": nil})
}

func (a *Api) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"OK": false, "payload": nil, "error": errors.New("wrong request format")})
		return
	}

	res, err := a.storage.DeleteById(oid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"OK": false, "payload": nil, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": true, "payload": res, "error": nil})
}
