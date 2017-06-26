package main

import (
	"encoding/json"

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
var mockData []byte

func init() {
	mockData = []byte(`[
	{
		"id": 1,
		"description": "test1",
		"like": {
			"count": 1,
			"name": [
				"first name",
				"peach",
				"pear"
			]
		}
	},
	{
		"id": 2,
		"description": "test2",
		"like": {
			"count": 1,
			"name": [
				"second name",
				"hello",
				"somkait"
			]
		}
	},
	{
		"id": 3,
		"description": "test3",
		"like": {
			"count": 1,
			"name": [
				"third name",
				"world",
				"pear"
			]
		}
	}
]`)
}

func main() {
	r := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/feed", func(c *gin.Context) {
		err := json.Unmarshal(mockData, &feed)
		if err != nil {
			panic(err)
		}
		c.JSON(200, feed)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
