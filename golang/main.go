package main

import "gopkg.in/gin-gonic/gin.v1"

type like struct {
	Count int      `json:"count"`
	Name  []string `json:"name"`
}

type data struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Like        like   `json:"like"`
}

var feed []data

func init() {
	feed = []data{
		data{
			ID:          1,
			Description: "test1",
			Like: like{
				Count: 1,
				Name:  []string{"first name", "peach", "pear"},
			},
		},
		data{
			ID:          2,
			Description: "test2",
			Like: like{
				Count: 1,
				Name:  []string{"second name", "hello", "somkait"},
			},
		},
		data{
			ID:          3,
			Description: "test3",
			Like: like{
				Count: 1,
				Name:  []string{"third name", "world", "pear"},
			},
		},
	}
}

func main() {
	feed[0].ID = 1
	// feed.Data[0].Description = "test"
	// feed.Data[0].Like.Count = 1
	// feed.Data[0].Like.Name = []string{"first name", "peach", "pear"}
	r := gin.Default()
	r.GET("/feed", func(c *gin.Context) {
		c.JSON(200, feed)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
