package auth

import "gorm.io/gorm"

type Usecase struct {
	db *gorm.DB
}

func NewUsecase(db *gorm.DB) *Usecase {
	return &Usecase{
		db: db,
	}
}
