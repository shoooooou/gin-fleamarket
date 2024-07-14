package main

import (
	"gin-fleamarket/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		items := []models.Item{
			{ID: 1, Name: "商品1", Price: 1000, Description: "商品１", SoldOut: false},
			{ID: 2, Name: "商品2", Price: 2000, Description: "商品2", SoldOut: true},
			{ID: 3, Name: "商品3", Price: 3000, Description: "商品3", SoldOut: false},
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
