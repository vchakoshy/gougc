package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vchakoshy/gougc/models"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB, router *gin.RouterGroup) *Module {
	db.AutoMigrate(&models.User{})

	m := NewDelivery(db)

	r := router.Group("/user")
	{
		r.GET("/", m.Index)

	}
	return &Module{
		db: db,
	}
}
