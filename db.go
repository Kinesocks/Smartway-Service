package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"db",
		"root",
		"root",
		"my_db",
		"5432",
		"disable",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	db.AutoMigrate(&Company{}, &Department{}, &Employee{}, &Passport{})

	return db
}
