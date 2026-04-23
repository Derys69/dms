package seeder

import (
	"log"

	"dms/internal/models"

	"gorm.io/gorm"
)

func RunSeeder(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Document{})

	users := []models.User{
		{Name: "Super Admin", Email: "admin@mail.com", Password: "password123", Role: "admin", Department: "Semua Departemen"},
		{Name: "Pengelola", Email: "pengelola@mail.com", Password: "password123", Role: "pengelola", Department: "IT"},
		{Name: "Staf IT", Email: "staf.it@mail.com", Password: "password123", Role: "pengguna_umum", Department: "IT"},
		{Name: "Staf HR", Email: "staf.hr@mail.com", Password: "password123", Role: "pengguna_umum", Department: "HR"},
	}

	for _, u := range users {
		db.Where("email = ?", u.Email).FirstOrCreate(&u)
	}

	var admin, managerIT, staffIT, staffHR models.User
	db.Where("email = ?", "admin@mail.com").First(&admin)
	db.Where("email = ?", "pengelola@mail.com").First(&managerIT)
	db.Where("email = ?", "staf.it@mail.com").First(&staffIT)
	db.Where("email = ?", "staf.hr@mail.com").First(&staffHR)

	documents := []models.Document{
		{Title: "Master Key System", Content: "Rahasia Tinggi.", OwnerID: admin.ID, Department: admin.Department},
		{Title: "Laporan Server IT Q1", Content: "Server berjalan normal.", OwnerID: managerIT.ID, Department: managerIT.Department},
		{Title: "Dokumentasi Kode API", Content: "Detail endpoint API.", OwnerID: staffIT.ID, Department: staffIT.Department},
		{Title: "Daftar Hadir Karyawan", Content: "Absensi harian.", OwnerID: staffHR.ID, Department: staffHR.Department},
	}

	for _, d := range documents {
		db.Where("title = ?", d.Title).FirstOrCreate(&d)
	}

	log.Println("Seeder berhasil dijalankan!")
}
