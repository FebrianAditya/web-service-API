package config

import (
	"fmt"
	"os"
	"../models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDBConnection connection to my database
func SetupDBConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env")
	}

	// process.env di GO contohnya kaya dibawah, kalau di JS kan process.env.DB_USER
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Cara connection dengan database di GO
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	// untuk baca dotEnv di Go pakai %s (contoh diatas)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connect database")
	}
	// db.AutoMigrate(structs.Person{})
	db.AutoMigrate()
	return db
}

// CloseDBConnection Closed database connection
func CloseDBConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from DB")
	}
	dbSQL.Close()
}