package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/controllers/catgoerycontroller"
	"github.com/taufik-hdyt/go-crud/controllers/foodcontroller"
	"github.com/taufik-hdyt/go-crud/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/api/foods", foodcontroller.GetAll)
	r.GET("/api/food/:id", foodcontroller.GetOne)
	r.POST("/api/food", foodcontroller.CreateRecipe)

	//catgeory
	r.GET("/api/categories", catgoerycontroller.GetCategories)
	r.POST("/api/categori", catgoerycontroller.CreateCategory)
	r.Run()
}
