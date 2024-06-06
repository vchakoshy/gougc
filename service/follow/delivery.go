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

// @Summary Follow User
// @Description Follow User
// @Accept  json
// @Produce  json
// @Tags Follow
// @Param data body FollowRequest true "data"
// @Success 200 {object} string
// @Router /follow/ [post]
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

// @Summary UnFollow User
// @Description UnFollow User
// @Accept  json
// @Produce  json
// @Tags Follow
// @Param data body FollowRequest true "data"
// @Success 200 {object} string
// @Router /follow/ [delete]
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
