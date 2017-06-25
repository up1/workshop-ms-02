# Feed

## GET /feed

Get result all feed.

===

### Example response

* Status: 200
* Content-Type: "application/json; charset=utf-8"

```json
[
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
]
```
