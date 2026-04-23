package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Ambil data dari .env isinya sesuai dengan db di pc
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	dbName := os.Getenv("DB_NAME")
	createQuery := fmt.Sprintf("CREATE DATABASE %s;", dbName)

	if err := db.Exec(createQuery).Error; err != nil {
		log.Println("Database mungkin sudah ada atau gagal dibuat:", err)
	} else {
		log.Println("Database", dbName, "berhasil dibuat!")
	}
}
