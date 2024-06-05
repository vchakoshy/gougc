package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB, router *gin.RouterGroup) *Module {
	// db.AutoMigrate(&models.Tag{})

	m := NewDelivery(db)

	r := router.Group("/auth")
	{
		r.GET("/", m.Index)

	}
	return &Module{
		db: db,
	}
}
