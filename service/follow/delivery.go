package follow

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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

// @Summary Followers list
// @Description Followers list
// @Accept  json
// @Produce  json
// @Tags Follow
// @Param user_id path int true "user_id"
// @Success 200 {object} []models.Follow
// @Router /follow/followers/{user_id} [get]
func (d Delivery) Followers(ctx *gin.Context) {
	uid := cast.ToUint(ctx.Param("user_id"))
	page := cast.ToInt(ctx.DefaultQuery("page", "0"))
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 10

	fl := d.usecase.Followers(uid, offset, 10)

	ctx.JSON(http.StatusOK, fl)
}

// @Summary Leaders list
// @Description Leaders list
// @Accept  json
// @Produce  json
// @Tags Follow
// @Param user_id path int true "user_id"
// @Success 200 {object} []models.Follow
// @Router /follow/leaders/{user_id} [get]
func (d Delivery) Leaders(ctx *gin.Context) {
	uid := cast.ToUint(ctx.Param("user_id"))
	page := cast.ToInt(ctx.DefaultQuery("page", "0"))
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 10

	fl := d.usecase.Leaders(uid, offset, 10)

	ctx.JSON(http.StatusOK, fl)
}
