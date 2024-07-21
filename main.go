package main

import (
	"fmt"
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	fmt.Println(os.Getenv("ENV"))
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "商品１", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "商品2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "商品3", SoldOut: false},
	}
	itemRepository := repositories.NewItemMemoryRepository(items)
	itemsService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemsService)
	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
