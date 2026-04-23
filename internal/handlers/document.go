package handlers

import (
	"net/http"

	"dms/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentHandler struct {
	DB *gorm.DB
}

func NewDocumentHandler(db *gorm.DB) *DocumentHandler {
	return &DocumentHandler{DB: db}
}

func (h *DocumentHandler) Create(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Judul dan isi dokumen wajib diisi"})
		return
	}

	valID, existsID := c.Get("userID")
	valDept, existsDept := c.Get("department")

	if !existsID || valID == nil || !existsDept || valDept == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Data identitas tidak lengkap dalam token"})
		return
	}

	var finalUserID uint
	switch v := valID.(type) {
	case uint:
		finalUserID = v
	case float64:
		finalUserID = uint(v)
	case int:
		finalUserID = uint(v)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Format User ID tidak valid"})
		return
	}

	finalDept, ok := valDept.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Format Department tidak valid"})
		return
	}

	newDoc := models.Document{
		Title:      req.Title,
		Content:    req.Content,
		OwnerID:    finalUserID,
		Department: finalDept,
	}

	if err := h.DB.Create(&newDoc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan dokumen"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Dokumen berhasil dibuat", "data": newDoc})
}

func (h *DocumentHandler) GetAll(c *gin.Context) {
	role, _ := c.Get("role")
	valID, _ := c.Get("userID")
	userDept, _ := c.Get("department")

	var docs []models.Document
	query := h.DB

	if role == "pengguna_umum" {
		query = query.Where("owner_id = ?", valID)
	} else if role == "pengelola" {
		query = query.Where("department = ?", userDept)
	}

	if err := query.Find(&docs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dokumen"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil mengambil data", "data": docs})
}

func (h *DocumentHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	var doc models.Document

	if err := h.DB.First(&doc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dokumen tidak ditemukan"})
		return
	}

	role, _ := c.Get("role")
	valID, _ := c.Get("userID")
	userDept, _ := c.Get("department")

	var currentUserID uint
	if v, ok := valID.(uint); ok {
		currentUserID = v
	} else if f, ok := valID.(float64); ok {
		currentUserID = uint(f)
	}

	if role == "pengguna_umum" && doc.OwnerID != currentUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak: Ini bukan dokumen Anda"})
		return
	}
	if role == "pengelola" && doc.Department != userDept.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak: Dokumen di luar departemen Anda"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": doc})
}

func (h *DocumentHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var doc models.Document

	if err := h.DB.First(&doc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dokumen tidak ditemukan"})
		return
	}

	role, _ := c.Get("role")
	userDept, _ := c.Get("department")

	if role == "pengguna_umum" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak: Pengguna umum tidak diizinkan menghapus"})
		return
	}

	if role == "pengelola" && doc.Department != userDept.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak: Departemen tidak sesuai"})
		return
	}

	if err := h.DB.Delete(&doc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus dokumen"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dokumen berhasil dihapus"})
}
