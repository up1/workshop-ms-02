package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

type SearchCommand struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.POST("/search", func(c *gin.Context) {
		var searchCmd SearchCommand
		c.BindJSON(&searchCmd)
		c.JSON(http.StatusOK, gin.H{"name": searchCmd.Name})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
