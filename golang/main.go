package main

import "gopkg.in/gin-gonic/gin.v1"

// You also can use a struct
var feed struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Like        struct {
		Count int      `json:"count"`
		Name  []string `json:"name"`
	} `json:"like"`
}

func main() {
	feed.ID = 1
	feed.Description = "test"
	feed.Like.Count = 1
	feed.Like.Name = []string{"first name", "peach", "pear"}
	r := gin.Default()
	r.GET("/feed", func(c *gin.Context) {
		c.JSON(200, feed)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
