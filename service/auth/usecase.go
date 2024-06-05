package auth

import (
	"errors"
	"time"

	"github.com/vchakoshy/gougc/models"
	"golang.org/x/crypto/bcrypt"
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

var ErrUserNotFound = errors.New("user not found")

func (u Usecase) Login(r LoginForm) (models.User, error) {
	var o models.User
	err := u.db.Model(&models.User{}).Where("username=?", r.Username).First(&o).Error
	if err != nil {
		return o, err
	}

	if !u.CheckPasswordHash(o.Password, r.Password) {
		return models.User{}, ErrUserNotFound
	}

	err = u.db.Model(&o).Update("last_login", time.Now()).Error

	return o, err

}

func (u Usecase) Register(r RegisterForm) (models.User, error) {
	hp, err := u.HashPassword(r.Password)
	if err != nil {
		return models.User{}, err
	}

	m := models.User{
		Username: r.Username,
		Password: hp,
	}

	err = u.db.Save(&m).Error

	return m, err
}

func (u *Usecase) HashPassword(p string) (string, error) {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func (u *Usecase) CheckPasswordHash(h string, p string) bool {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	// nil means it is a match
	return err == nil
}
