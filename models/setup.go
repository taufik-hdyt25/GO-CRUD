package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "host=ep-fragrant-bush-09941286-pooler.us-east-2.aws.neon.tech user=taufik-hdyt password=QNp8YA2bSzse dbname=CulinaryCompanionDB port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Food{}, &Category{})
	DB = database
}
