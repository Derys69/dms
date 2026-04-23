package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")

	AuthRoutes(api, db)
	DocumentRoutes(api, db)

	return r
}
