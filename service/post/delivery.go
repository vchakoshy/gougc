package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Delivery struct {
	usecase *Usecase
}

func NewDelivery(db *gorm.DB) *Delivery {
	ctrl := &Delivery{
		usecase: NewUsecase(db),
	}

	return ctrl
}

func (Delivery) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
