package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"dms/internal/dto"
	"dms/internal/models"
	"dms/internal/utils"

	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

var googleOauthConfig = &oauth2.Config{
	// Ganti dengan kredensial Google OAuth yang dibuat dan masukan di .env
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:3005/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	var user models.User
	if err := h.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Role, user.Department)
	c.JSON(http.StatusOK, dto.LoginResponse{
		Message: "Login Berhasil",
		Token:   token,
		User:    dto.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Role: user.Role, Department: user.Department},
	})
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) GoogleLoginCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Gagal tukar kode"})
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Gagal ambil profil"})
		return
	}
	defer resp.Body.Close()

	var googleUser struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	json.NewDecoder(resp.Body).Decode(&googleUser)

	var user models.User
	if err := h.DB.Where("google_id = ?", googleUser.ID).First(&user).Error; err != nil {
		user = models.User{
			Name:       googleUser.Name,
			Email:      googleUser.Email,
			GoogleID:   googleUser.ID,
			Role:       "pengguna_umum",
			Department: "Umum",
		}
		h.DB.Create(&user)
	}

	jwtToken, _ := utils.GenerateToken(user.ID, user.Role, user.Department)

	// ARAHKAN KE UI (INDEX.HTML)
	// Sesuaikan http://127.0.0.1:5500 dengan alamat yang muncul di browser saat klik 'Open with Live Server' jika menggunakkan vscode
	c.Redirect(http.StatusTemporaryRedirect, "http://127.0.0.1:5500/UI/index.html?token="+jwtToken)
}
