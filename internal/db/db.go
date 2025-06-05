package db

import (
	"Tasks/internal/taskService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return nil, err
	}

	if err := DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
		return nil, err
	}

	log.Println("Database migrated successfully")
	return DB, nil
}
