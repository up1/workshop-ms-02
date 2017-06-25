package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-gin-prometheus"
)

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
	r := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/feed", func(c *gin.Context) {
		c.JSON(200, feed)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
