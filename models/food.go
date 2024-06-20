package models

type Food struct {
	Id        int64    `gorm:"primaryKey" json:"id"`
	Name      string   `gorm:"type:varchar(300)" json:"name"`
	Deskripsi string   `gorm:"type:text" json:"deskripsi"`
	Bahan     []string `gorm:"type:text[]" json:"bahan"`
	Step      []string `gorm:"type:text[]" json:"step"`
	Image     string   `gorm:"type:varchar(300)" json:"image"`
}
