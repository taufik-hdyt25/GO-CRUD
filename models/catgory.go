package models

type Category struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name"`
}
