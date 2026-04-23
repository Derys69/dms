package routes

import (
	"dms/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	authHandler := handlers.NewAuthHandler(db)

	auth := rg.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.GET("/google/login", authHandler.GoogleLogin)
		auth.GET("/google/callback", authHandler.GoogleLoginCallback)
	}
}
