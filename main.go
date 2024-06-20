package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/controllers/foodcontroller"
	"github.com/taufik-hdyt/go-crud/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/api/foods", foodcontroller.GetAll)
	r.GET("/api/food/:id", foodcontroller.GetOne)

	r.Run()
}
