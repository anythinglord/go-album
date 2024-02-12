package main

import (
	"github.com/gin-gonic/gin"
)

type Album struct {
	id     string  `json:"id"`
	title  string  `json:"title"`
	artist string  `json:"artist"`
	price  float64 `json:"price"`
}

var albums = []Album{
	{id: "1", title: "Blue Train", artist: "Jhon Coltrane", price: 56.95},
	{id: "2", title: "Jeru", artist: "Gerry Mulligan", price: 56.95},
	{id: "3 ", title: "Clifford Brown", artist: "Sarah Vaughan", price: 56.95},
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.Run("localhost:8080")
}
