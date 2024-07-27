package main

import (
	"fmt"
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"

	// "gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	fmt.Println(os.Getenv("ENV"))
	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "商品１", SoldOut: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "商品2", SoldOut: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "商品3", SoldOut: false},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)

	itemsService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemsService)
	r := gin.Default()
	itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
