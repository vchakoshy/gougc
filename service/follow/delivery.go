package follow

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Delivery struct {
	usecase *Usecase
}

func NewDelivery(db *gorm.DB, uc *Usecase) *Delivery {
	ctrl := &Delivery{
		usecase: uc,
	}

	return ctrl
}

func (d Delivery) Follow(ctx *gin.Context) {
	var r FollowRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := d.usecase.Follow(r.LeaderID, r.FollowerID); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}

func (d Delivery) UnFollow(ctx *gin.Context) {
	var r FollowRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := d.usecase.UnFollow(r.LeaderID, r.FollowerID); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}
