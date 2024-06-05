package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vchakoshy/gougc/models"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB) *Module {
	db.AutoMigrate(&models.User{})

	return &Module{
		db: db,
	}
}

func (m *Module) SetupRoutes(router *gin.RouterGroup) {
	d := NewDelivery(m.db)
	r := router.Group("/auth")
	{
		r.POST("/register/", d.Register)
		r.POST("/login/", d.Login)
	}
}
