package post

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB) *Module {
	return &Module{
		db: db,
	}
}

func (m *Module) SetupRoutes(router *gin.RouterGroup) {
	d := NewDelivery(m.db)
	r := router.Group("/post")
	{
		r.GET("/", d.Index)
	}
}
