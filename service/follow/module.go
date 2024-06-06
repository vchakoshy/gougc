package follow

import (
	"github.com/gin-gonic/gin"
	"github.com/vchakoshy/gougc/models"
	"gorm.io/gorm"
)

type Module struct {
	db      *gorm.DB
	Usecase *Usecase
}

func NewModule(db *gorm.DB) *Module {
	db.AutoMigrate(&models.Follow{})

	return &Module{
		db:      db,
		Usecase: NewUsecase(db),
	}
}

func (m *Module) SetupRoutes(router *gin.RouterGroup) {
	d := NewDelivery(m.db, m.Usecase)
	r := router.Group("/follow")
	{
		r.POST("/", d.Follow)
		r.DELETE("/", d.UnFollow)

		r.GET("/followers/:user_id", d.Followers)
		r.GET("/leaders/:user_id", d.Leaders)
	}
}
