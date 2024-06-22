package models

import (
	"mime/multipart"

	"github.com/lib/pq"
)

type Food struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	CategoryID  uint           `json:"category_id"`
	Steps       pq.StringArray `gorm:"type:text[]" json:"steps"`
	Ingredients pq.StringArray `gorm:"type:text[]" json:"ingredients"`
	Category    Category       `json:"category"`
}
type CreateRecipeInput struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	Ingredients string                `form:"ingredients[]" binding:"required"`
	Steps       string                `form:"steps[]" binding:"required"`
	CategoryID  uint                  `form:"category_id" binding:"required"`
}

type Foods struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	Ingredients pq.StringArray `gorm:"type:text[]" json:"ingredients"`
	Steps       pq.StringArray `gorm:"type:text[]" json:"steps"`
	CategoryID  uint           `json:"category_id"`
}
