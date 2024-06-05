package follow

import (
	"github.com/gin-gonic/gin"
	"github.com/vchakoshy/gougc/models"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB, router *gin.RouterGroup) *Module {
	db.AutoMigrate(&models.Follow{})

	m := NewDelivery(db)

	r := router.Group("/follow")
	{
		r.POST("/", m.Follow)
		r.DELETE("/", m.UnFollow)

	}
	return &Module{
		db: db,
	}
}
