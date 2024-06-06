package follow

import (
	"github.com/vchakoshy/gougc/models"
	"gorm.io/gorm"
)

type Usecase struct {
	db *gorm.DB
}

func NewUsecase(db *gorm.DB) *Usecase {
	return &Usecase{
		db: db,
	}
}

func (u Usecase) Follow(leaderID, followerID uint) error {
	o := models.Follow{
		LeaderID:   leaderID,
		FollowerID: followerID,
	}

	return u.db.Create(&o).Error
}

func (u Usecase) UnFollow(leaderID, followerID uint) error {
	var o models.Follow
	err := u.db.Model(&models.Follow{}).
		Where(models.Follow{LeaderID: leaderID, FollowerID: followerID}).
		First(&o).Error
	if err != nil {
		return err
	}

	return u.db.Delete(&o).Error
}

func (u Usecase) IsFollow(followerID, leaderID uint) bool {
	var count int64
	u.db.Model(&models.Follow{}).Where(models.Follow{
		LeaderID: leaderID, FollowerID: followerID,
	}).Count(&count)

	return count > 0
}

func (u Usecase) Followers(leaderID uint, offset, limit int) []models.Follow {
	var fl []models.Follow
	u.db.Model(&models.Follow{}).
		Where(models.Follow{LeaderID: leaderID}).
		Offset(offset).
		Limit(limit).
		Find(&fl)

	return fl
}

func (u Usecase) Leaders(followerID uint, offset, limit int) []models.Follow {
	var fl []models.Follow
	u.db.Model(&models.Follow{}).
		Where(models.Follow{FollowerID: followerID}).
		Offset(offset).
		Limit(limit).
		Find(&fl)

	return fl
}
