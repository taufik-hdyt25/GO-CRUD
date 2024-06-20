package foodcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/models"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	var foods []models.Food
	var totalItems int64

	// Get page and limit from query parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get total count of items
	models.DB.Model(&models.Food{}).Count(&totalItems)

	// Fetch items with pagination
	models.DB.Offset(offset).Limit(limit).Find(&foods)

	// Calculate next page
	var nextPage *int
	if int64(offset+limit) < totalItems {
		next := page + 1
		nextPage = &next
	} else {
		nextPage = nil
	}

	// Respond with JSON
	c.JSON(http.StatusOK, gin.H{
		"foods":     foods,
		"page":      page,
		"limit":     limit,
		"total":     totalItems,
		"next_page": nextPage,
	})
}

func GetOne(c *gin.Context) {
	var food models.Food
	id := c.Param("id")
	if err := models.DB.First(&food, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak di temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"food": food})
}
