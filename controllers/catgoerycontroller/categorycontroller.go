package catgoerycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/models"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	models.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&category)
	c.JSON(http.StatusOK, category)
}
