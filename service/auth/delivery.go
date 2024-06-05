package auth

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

func (d Delivery) Register(ctx *gin.Context) {
	var r RegisterForm
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	o, err := d.usecase.Register(r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, o)
}

func (d Delivery) Login(ctx *gin.Context) {
	var r LoginForm
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	o, err := d.usecase.Login(r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, o)
}
