package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Koneksi ke db bawaan 'postgres' untuk membuat database baru
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Langsung eksekusi buat database
	db.Exec("CREATE DATABASE server_db;")
}
