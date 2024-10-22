package models

import (
	"fmt"
	"log"
	"santrikoding/backend-api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t&loc=%s",
		config.AppConfig.GetString("db.user"),
		config.AppConfig.GetString("db.password"),
		config.AppConfig.GetString("db.host"),
		config.AppConfig.GetInt("db.port"),
		config.AppConfig.GetString("db.name"),
		config.AppConfig.GetBool("db.parse_time"),
		config.AppConfig.GetString("db.loc"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	DB = database

	if DB == nil {
		log.Fatal("Database connection is nil")
	}

	if err := DB.AutoMigrate(&Post{}, &User{}); err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}

	log.Println("Database connected and tables migrated successfully")
}
