package attachment

import (
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

func (Delivery) Upload(ctx *gin.Context) {

}
