package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type container struct {
	ID	string `json:"id"`
	Name	string `json:"name"`
	IP	string `json:"ip"`
	Tag	string `json:"tag"`
}

var containers = []container{}

func main() {
	router := gin.Default()
	router.GET("/containers", getContainers)
	router.POST("/containers", postContainers)

	router.Run("localhost:8080")
}

func getContainers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, containers)
}

func postContainers(c *gin.Context) {
	var newContainer container

	if err := c.BindJSON(&newContainer); err != nil {
		return
	}

	containers = append(containers, newContainer)
	c.IndentedJSON(http.StatusCreated, newContainer)
}
