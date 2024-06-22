package foodcontroller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/taufik-hdyt/go-crud/models"
	"github.com/taufik-hdyt/go-crud/services"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	var foods []models.Food
	var totalItems int64

	// Default values for page and pageSize
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Query to count total items
	models.DB.Model(&models.Food{}).Count(&totalItems)

	// Query to get paginated results
	models.DB.Limit(pageSize).Offset(offset).Find(&foods)

	// Prepare response with pagination metadata
	response := gin.H{
		"foods": foods,
		"pagination": gin.H{
			"page":       page,
			"pageSize":   pageSize,
			"totalItems": totalItems,
			"totalPages": (totalItems + int64(pageSize) - 1) / int64(pageSize), // Calculate total pages
		},
	}

	c.JSON(http.StatusOK, response)
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

func CreateRecipe(c *gin.Context) {
	var input models.CreateRecipeInput

	// Bind form data to the struct
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if input.Name == "" || input.Description == "" || len(input.Ingredients) == 0 || len(input.Steps) == 0 || input.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Handle image upload
	imageFile, err := input.Image.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process uploaded image"})
		return
	}
	defer imageFile.Close()

	// Save the uploaded file locally
	var imageUrl string
	imagePath := fmt.Sprintf("./temp/%s", input.Image.Filename)
	if err := c.SaveUploadedFile(input.Image, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded image"})
		return
	}

	// Upload the image to Cloudinary (or any other service)
	cloudinaryService := services.NewCloudinaryService("cloudinary://866624875478827:-iE68VPffoTGtfUUS1ItgtDRI_0@doushe6hn", "taufikhdyt", "866624875478827", "-iE68VPffoTGtfUUS1ItgtDRI_0")
	imageUrl, err = cloudinaryService.UploadImage(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Optionally, delete the local file after upload
	defer os.Remove(imagePath)

	// Create Recipe
	recipe := models.Foods{
		Name:        input.Name,
		Description: input.Description,
		Image:       imageUrl,
		Ingredients: pq.StringArray{input.Ingredients},
		Steps:       pq.StringArray{input.Steps},
		CategoryID:  input.CategoryID,
	}

	// Save recipe to database
	if err := models.DB.Create(&recipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_create_food": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}
