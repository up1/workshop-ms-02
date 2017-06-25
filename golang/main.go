package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	r := gin.Default()
	r.GET("/feed", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id":          1,
			"description": "test",
			"like":        1,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
