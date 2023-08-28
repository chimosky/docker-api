package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type container struct {
	Name	string `json:"name"`
	Tag	string `json:"tag"`
}

var containers []containerDetails

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
	var details containerDetails

	if err := c.BindJSON(&newContainer); err != nil {
		return
	}

        if tag := newContainer.Tag; tag != "latest" {
		details = Run(newContainer.Name, tag)
	} else {
		details = Run(newContainer.Name, "latest")
	}

	containers = append(containers, details)
	c.IndentedJSON(http.StatusCreated, details)
}
