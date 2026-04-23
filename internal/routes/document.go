package routes

import (
	"dms/internal/handlers"
	"dms/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DocumentRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewDocumentHandler(db)
	docs := rg.Group("/documents", middleware.AuthMiddleware())
	{
		docs.POST("/", h.Create)
		docs.GET("/", h.GetAll)
		docs.GET("/:id", h.GetByID)
		docs.DELETE("/:id", h.Delete)
	}
}
